package main

import (
	"archive/zip"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var returnString = "Default Message. If you are seeing this, Corey, Trevor... You fucked up!"
	var sendMessage = false

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == commandPrefix+"get event" {
		logger.Info().Msg("Message Create Event - Get Event - Triggered")
		var nextEvent, err = lanopsAPI.GetNextEvent()
		if err != nil {
			returnString = fmt.Sprintf("Something went wrong: %s", err)
		} else {
			returnString = formatNextEventMessage(nextEvent)
		}
		sendMessage = true
	}

	if m.Content == commandPrefix+"get dates" {
		logger.Info().Msg("Message Create Event - Get Dates - Triggered")
		var upcomingEvents, err = lanopsAPI.GetUpcomingEvents()
		if err != nil {
			returnString = fmt.Sprintf("Something went wrong: %s", err)
		} else {
			returnString = formatUpcomingEventDatesMessage(upcomingEvents)
		}
		sendMessage = true
	}

	if m.Content == commandPrefix+"get attendees" {
		logger.Info().Msg("Message Create Event - Get Attendees - Triggered")
		var participants, err = lanopsAPI.GetNextEventParticipants()
		if err != nil {
			returnString = fmt.Sprintf("Something went wrong: %s", err)
		} else {
			returnString = formatEventParticipantsMessage(participants)
		}
		sendMessage = true
	}

	if discordJukeboxControlEnabled {
		if m.Content == commandPrefix+"jb current" {
			logger.Info().Msg("Message Create Event - Jukebox Currently playing - Triggered")
			returnString = jukeboxAPI.GetCurrentTrack()
			sendMessage = true
		} else if slices.Contains(m.Member.Roles, discordJukeBoxControlRoleID) {
			if strings.HasPrefix(m.Content, commandPrefix+"jb") {
				logger.Info().Msg("Message Create Event - Jukebox Control - Triggered")
				jukeboxCommand := strings.Split(m.Content, " ")
				returnString = jukeboxAPI.Control(jukeboxCommand[1])
				sendMessage = true
			}
		}
	}

	// Image Archiver
	if m.Content == commandPrefix+"image archive" {
		logger.Info().Msg("Message Create Event - Image Archive - Triggered")
		returnString = "Archiving Channel"

		filesToUpload := false
		var lastMessageID string
		for {
			messages, err := dg.ChannelMessages(m.ChannelID, 100, lastMessageID, "", "")
			if err != nil {
				fmt.Println("Error fetching messages:", err)
				return
			}
			if len(messages) == 0 {
				break // No more messages
			}
			for _, msg := range messages {
				if len(msg.Attachments) > 0 {
					for _, attachment := range msg.Attachments {
						if !isImage(attachment.Filename) {
							continue // Skip non-image files
						}
						fmt.Printf("Attachment: %s (URL: %s)\n", attachment.Filename, attachment.URL)
						err := downloadFile(attachment.URL, attachment.Filename)
						if err != nil {
							fmt.Println("Failed to download:", err)
						}
						filesToUpload = true
					}
				}
			}
			lastMessageID = messages[len(messages)-1].ID // Set the last message ID for pagination
		}
		if filesToUpload {
			channel, err := dg.Channel(m.ChannelID)
			if err != nil {
				fmt.Println("Error fetching channel:", err)
				return
			}
			zipPath := fmt.Sprintf("/tmp/%s-%s-attachments.zip", time.Now().Format("2006-01-02_15-04-05"), channel.Name)
			err = zipDirectory("tmp/attachments", zipPath)
			if err != nil {
				fmt.Println("Failed to download:", err)
			}
			ctx := context.Background()

			googleCreds, err := os.ReadFile("googlekey.json")
			if err != nil {
				fmt.Printf("Unable to read credentials file: %v\n", err)
				os.Exit(1)
			}

			// Authenticate using the service account
			googleConfig, err := google.JWTConfigFromJSON(googleCreds, drive.DriveFileScope)
			if err != nil {
				fmt.Printf("Unable to parse client secret file to config: %v\n", err)
				os.Exit(1)
			}

			googleClient := googleConfig.Client(ctx)

			// Create Drive service
			googleSrv, err := drive.NewService(ctx, option.WithHTTPClient(googleClient))
			if err != nil {
				fmt.Printf("Unable to retrieve Drive client: %v\n", err)
				os.Exit(1)
			}

			parentFolderID := googleDriveUploadDirectoryID

			fileID, err := uploadToDrive(googleSrv, zipPath, parentFolderID)
			if err != nil {
				fmt.Printf("Error uploading %s: %v\n", zipPath, err)
			}

			returnString = fmt.Sprintf("Download link: https://drive.google.com/file/d/%s/view", fileID)

			if err != nil {
				fmt.Println("Error uploading zip:", err)
			} else {
				fmt.Printf("Successfully uploaded %s to Drive!", zipPath)
			}
			err = os.RemoveAll("tmp")
			if err != nil {
				fmt.Println("Error deleting tmp folder:", err)
			} else {
				fmt.Println("Deleted tmp folder.")
			}
		}

		sendMessage = true
	}

	// Memes
	if m.Author.ID == memeNameChangerUserID {
		userNames := []string{
			"Dumbbell Chrome Remover",
			"Jay2Win",
			"Perry",
			"Frank Reynolds",
			"Scraninator",
			"Lord Scranian",
			"Eddy Hall",
			"Scran Master",
			"Scran2D2",
			"Bruce Scranner",
			"Scranuel Jackson",
			"Protein Baggins",
			"Scran Solo",
			"Jason Gainham",
			"Captain Ameri-scran",
			"Whoopi Swoleberg",
			"Scranakin Skywalker",
			"Obi-Wan Scranobi",
			"The Swole Ranger",
			"Gains Bond",
		}
		randomIndex := rand.Intn(len(userNames))
		randomString := userNames[randomIndex]

		err := dg.GuildMemberNickname(discordGuildID, memeNameChangerUserID, randomString)
		if err != nil {
			fmt.Println("error changing nickname,", err)
			return
		}
	}

	// Return the Message
	if sendMessage {
		s.ChannelMessageSend(m.ChannelID, returnString)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "Lan Organization")
}

func isImage(filename string) bool {
	lower := strings.ToLower(filename)
	return strings.HasSuffix(lower, ".jpg") ||
		strings.HasSuffix(lower, ".jpeg") ||
		strings.HasSuffix(lower, ".png") ||
		strings.HasSuffix(lower, ".gif") ||
		strings.HasSuffix(lower, ".webp")
}

func downloadFile(url, filename string) error {
	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := os.Stat("tmp/attachments"); os.IsNotExist(err) {
		return os.MkdirAll("tmp/attachments", 0755) // Creates all needed folders with permissions
	}

	// Create the file
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return err
	}
	out, err := os.Create("tmp/attachments/" + hex.EncodeToString(bytes) + "-" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func zipDirectory(sourceDir, zipFileName string) error {
	// Create the zip file
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through every file in the folder
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore directories, only files
		if info.IsDir() {
			return nil
		}

		// Create zip header
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate // Compression method

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Open the file and copy it into the zip
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})
}

func uploadToDrive(srv *drive.Service, filePath string, parentFolderID string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	f := &drive.File{
		Name:    filepath.Base(filePath),
		Parents: []string{parentFolderID},
	}

	createdFile, err := srv.Files.Create(f).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("could not upload file: %w", err)
	}

	fmt.Printf("Uploaded: %s\n", filepath.Base(filePath))
	return createdFile.Id, nil
}
