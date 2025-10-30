package mediaarchiver

import (
	"fmt"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/utils"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {

	if slices.Contains(m.Member.Roles, cfg.Discord.MediaArchiverRoleId) {
		var returnString string
		msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Archive Media - Triggered", Level: "INFO"}

		returnString = "Archiving Channel Media!"
		var err error
		daysRangeInt := 0
		archiveCommand := strings.Split(m.Content, " ")
		if len(archiveCommand) == 3 {
			daysRangeInt, err = strconv.Atoi(args[0])
		}
		if err != nil {
			returnString = "Invalid Days Parameter"
		} else {
			go archiveChannelMedia(s, m, daysRangeInt, cfg, msgCh)
		}

		s.ChannelMessageSend(m.ChannelID, returnString)
	}
}

func archiveChannelMedia(s *discordgo.Session, m *discordgo.MessageCreate, dateRangeDays int, cfg config.Config, msgCh chan<- channels.MsgCh) {
	var lastMessageID string
	filesToUpload := false
	downloadCount := 0
	parentFolderID := cfg.MediaArchiver.GoogleDriveUploadDirId

	randomDirectoryName := fmt.Sprintf(
		"%s-%s-%s",
		m.ChannelID,
		m.Author.ID,
		m.Message.Timestamp.Format("2006-01-02-15-04-05"),
	)

	archiveChannelMediaTmpDirPath := strings.TrimSuffix(cfg.MediaArchiver.TmpUploadDirPath, "/") + "/" + randomDirectoryName
	downloadPath := archiveChannelMediaTmpDirPath + "/attachments"

	archiveTenseMessage := fmt.Sprintf("last %s days of media", strconv.Itoa(dateRangeDays))
	if dateRangeDays == 0 {
		archiveTenseMessage = "ALL media"
	}

	msgCh <- channels.MsgCh{Err: nil, Message: fmt.Sprintf("Attempting to archive %s in channel", archiveTenseMessage), Level: "INFO"}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Attempting to archive %s in channel", archiveTenseMessage))

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Error fetching channel", Level: "ERROR"}
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error fetching channel: %s", err))
		return
	}

	for {
		dateRangeExceededSkip := false
		messages, err := s.ChannelMessages(m.ChannelID, 100, lastMessageID, "", "")
		if err != nil {
			msgCh <- channels.MsgCh{Err: err, Message: "Error fetching messages", Level: "ERROR"}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error fetching messages: %s", err))
			return
		}
		if len(messages) == 0 {
			break // No more messages
		}
		for _, msg := range messages {
			if dateRangeDays != 0 && msg.Timestamp.Before(time.Now().AddDate(0, 0, -dateRangeDays)) {
				dateRangeExceededSkip = true
				break
			}
			if len(msg.Attachments) > 0 {
				for _, attachment := range msg.Attachments {
					if !utils.IsMedia(attachment.Filename) {
						continue
					}
					msgCh <- channels.MsgCh{Err: nil, Message: fmt.Sprintf("Attachment: %s (URL: %s)\n", attachment.Filename, attachment.URL), Level: "INFO"}

					randomFileName, err := utils.RandomString(8)
					if err != nil {
						msgCh <- channels.MsgCh{Err: err, Message: "Error generating random name for file", Level: "ERROR"}
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error generating random name for file: %s", err))
						continue
					}
					err = utils.DownloadFile(
						utils.DownloadFileParams{
							URL:          attachment.URL,
							Filename:     randomFileName + attachment.Filename,
							DownloadPath: downloadPath,
						})
					if err != nil {
						msgCh <- channels.MsgCh{Err: err, Message: "Failed to download media", Level: "ERROR"}
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Failed to download media: %s", err))
						return
					} else {
						filesToUpload = true
						downloadCount++
					}
				}
			}
		}
		if dateRangeExceededSkip {
			break
		}
		lastMessageID = messages[len(messages)-1].ID // Set the last message ID for pagination
	}

	time.Sleep(2000 * time.Millisecond) // Wait to get around weird bugs with the downloads

	msgCh <- channels.MsgCh{Err: nil, Message: fmt.Sprintf("Media to upload: %s", strconv.Itoa(downloadCount)), Level: "INFO"}
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Media to upload: %s", strconv.Itoa(downloadCount)))

	if filesToUpload {
		zipPath := fmt.Sprintf(
			"%s/%s-%s-attachments.zip",
			archiveChannelMediaTmpDirPath,
			time.Now().Format("2006-01-02"),
			channel.Name)
		err = utils.ZipDirectory(downloadPath, zipPath)
		if err != nil {
			msgCh <- channels.MsgCh{Err: err, Message: "Failed to zip attachments", Level: "ERROR"}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Failed to zip attachments: %s", err))
			go cleanup(s, m, archiveChannelMediaTmpDirPath, msgCh)
			return
		}

		fileID, err := utils.UploadToDrive(zipPath, parentFolderID, cfg.MediaArchiver.GoogleDriveKey)
		if err != nil {
			msgCh <- channels.MsgCh{Err: err, Message: "Error uploading", Level: "ERROR"}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error uploading: %s", err))
			go cleanup(s, m, archiveChannelMediaTmpDirPath, msgCh)
			return
		}

		msgCh <- channels.MsgCh{Err: nil, Message: fmt.Sprintf("Uploaded: %s\n", filepath.Base(zipPath)), Level: "INFO"}

		if err := utils.DmUser(s, m.Author.ID, fmt.Sprintf("Download link: https://drive.google.com/file/d/%s/view", fileID)); err != nil {
			msgCh <- channels.MsgCh{Err: err, Message: "Error sending DM to User", Level: "ERROR"}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error sending DM to User: %s", err))
			go cleanup(s, m, archiveChannelMediaTmpDirPath, msgCh)
		} else {
			s.ChannelMessageSend(m.ChannelID, "You have been sent a link to the download")
		}
		go cleanup(s, m, archiveChannelMediaTmpDirPath, msgCh)
	}
}

func cleanup(s *discordgo.Session, m *discordgo.MessageCreate, archiveChannelMediaTmpDirPath string, msgCh chan<- channels.MsgCh) {
	if err := utils.DeleteDir(archiveChannelMediaTmpDirPath); err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Error deleting tmp folder", Level: "ERROR"}
		s.ChannelMessageSend(m.ChannelID, "Error deleting tmp folder")
	}
}
