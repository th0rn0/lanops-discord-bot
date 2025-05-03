package main

import (
	"fmt"
	"lanops/discord-bot/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func archiveChannelMedia(m *discordgo.MessageCreate, dateRangeDays int) {
	var lastMessageID string
	filesToUpload := false
	downloadCount := 0
	parentFolderID := os.Getenv("GOOGLE_DRIVE_UPLOAD_DIRECTORY_ID")

	randomDirectoryName := fmt.Sprintf(
		"%s-%s-%s",
		m.ChannelID,
		m.Author.ID,
		m.Message.Timestamp.Format("2006-01-02-15-04-05"),
	)

	archiveChannelMediaTmpDirPath = strings.TrimSuffix(os.Getenv("ARCHIVE_CHANNEL_MEDIA_TMP_DIR_PATH"), "/") + "/" + randomDirectoryName
	downloadPath := archiveChannelMediaTmpDirPath + "/attachments"

	archiveTenseMessage := fmt.Sprintf("last %s days of media", strconv.Itoa(dateRangeDays))
	if dateRangeDays == 0 {
		archiveTenseMessage = "ALL media"
	}
	logger.Info().Msg(fmt.Sprintf("Attempting to archive %s in channel", archiveTenseMessage))
	addMessageToQueue(m.ChannelID, fmt.Sprintf("Attempting to archive %s in channel", archiveTenseMessage))

	channel, err := dg.Channel(m.ChannelID)
	if err != nil {
		logger.Error().Err(err).Msg("Error fetching channel")
		go addMessageToQueue(m.ChannelID, fmt.Sprintf("Error fetching channel: %s", err))
		return
	}

	for {
		dateRangeExceededSkip := false
		messages, err := dg.ChannelMessages(m.ChannelID, 100, lastMessageID, "", "")
		if err != nil {
			logger.Error().Err(err).Msg("Error fetching messages")
			go addMessageToQueue(m.ChannelID, fmt.Sprintf("Error fetching messages: %s", err))
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
					logger.Info().Msg(fmt.Sprintf("Attachment: %s (URL: %s)\n", attachment.Filename, attachment.URL))

					randomFileName, err := utils.RandomString(8)
					if err != nil {
						logger.Error().Err(err).Msg("Error generating random name for file")
						go addMessageToQueue(m.ChannelID, fmt.Sprintf("Error generating random name for file: %s", err))
						continue
					}
					err = utils.DownloadFile(
						utils.DownloadFileParams{
							URL:          attachment.URL,
							Filename:     randomFileName + attachment.Filename,
							DownloadPath: downloadPath,
						})
					if err != nil {
						logger.Error().Err(err).Msg("Failed to download")
						go addMessageToQueue(m.ChannelID, fmt.Sprintf("Failed to download media: %s", err))
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

	logger.Info().Msg(fmt.Sprintf("Media to upload: %s", strconv.Itoa(downloadCount)))
	addMessageToQueue(m.ChannelID, fmt.Sprintf("Media to upload: %s", strconv.Itoa(downloadCount)))

	if filesToUpload {
		zipPath := fmt.Sprintf(
			"%s/%s-%s-attachments.zip",
			archiveChannelMediaTmpDirPath,
			time.Now().Format("2006-01-02"),
			channel.Name)
		err = utils.ZipDirectory(downloadPath, zipPath)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to zip attachments")
			go addMessageToQueue(m.ChannelID, fmt.Sprintf("Failed to zip attachments: %s", err))
			go cleanup(archiveChannelMediaTmpDirPath, m.ChannelID)
			return
		}

		fileID, err := utils.UploadToDrive(zipPath, parentFolderID)
		if err != nil {
			logger.Error().Err(err).Msg("Error uploading")
			go addMessageToQueue(m.ChannelID, fmt.Sprintf("Error uploading: %s", err))
			go cleanup(archiveChannelMediaTmpDirPath, m.ChannelID)
			return
		}
		logger.Info().Msg(fmt.Sprintf("Uploaded: %s\n", filepath.Base(zipPath)))

		if err := utils.DmUser(dg, m.Author.ID, fmt.Sprintf("Download link: https://drive.google.com/file/d/%s/view", fileID)); err != nil {
			logger.Error().Err(err).Msg("Error sending DM to User")
			go addMessageToQueue(m.ChannelID, fmt.Sprintf("Error sending DM of link to User: %s", err))
		} else {
			go addMessageToQueue(m.ChannelID, "You have been sent a link to the download")
		}
		go cleanup(archiveChannelMediaTmpDirPath, m.ChannelID)
	}
}

func cleanup(archiveChannelMediaTmpDirPath string, channelID string) {
	if err := utils.DeleteDir(archiveChannelMediaTmpDirPath); err != nil {
		logger.Error().Err(err).Msg("Error deleting tmp folder")
		addMessageToQueue(channelID, fmt.Sprintf("Error deleting tmp directory: %s", archiveChannelMediaTmpDirPath))
	}
}
