package main

import (
	"github.com/bwmarrin/discordgo"
)

var webhookEnabled = false

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var returnString = "Default Message. If you are seeing this, Corey, Trevor... You fucked up!"
	var sendMessage = false

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if m.Content == commandPrefix+"get event" {
	// 	logger.Info().Msg("Message Create Event - Get Event - Triggered")
	// 	var nextEvent, err = lanopsAPI.GetNextEvent()
	// 	if err != nil {
	// 		returnString = fmt.Sprintf("Something went wrong: %s", err)
	// 	} else {
	// 		returnString = formatNextEventMessage(nextEvent)
	// 	}
	// 	sendMessage = true
	// }

	// if m.Content == commandPrefix+"get dates" {
	// 	logger.Info().Msg("Message Create Event - Get Dates - Triggered")
	// 	var upcomingEvents, err = lanopsAPI.GetUpcomingEvents()
	// 	if err != nil {
	// 		returnString = fmt.Sprintf("Something went wrong: %s", err)
	// 	} else {
	// 		returnString = formatUpcomingEventDatesMessage(upcomingEvents)
	// 	}
	// 	sendMessage = true
	// }

	// if m.Content == commandPrefix+"get attendees" {
	// 	logger.Info().Msg("Message Create Event - Get Attendees - Triggered")
	// 	var participants, err = lanopsAPI.GetNextEventParticipants()
	// 	if err != nil {
	// 		returnString = fmt.Sprintf("Something went wrong: %s", err)
	// 	} else {
	// 		returnString = formatEventParticipantsMessage(participants)
	// 	}
	// 	sendMessage = true
	// }

	// Media Archiver
	// if slices.Contains(m.Member.Roles, archiveChannelMediaRoleID) {
	// 	if strings.HasPrefix(m.Content, commandPrefix+"media archive") {
	// 		logger.Info().Msg("Message Create Event - Image Archive - Triggered")
	// 		returnString = "Archiving Channel Media!"
	// 		var err error
	// 		daysRangeInt := 0
	// 		archiveCommand := strings.Split(m.Content, " ")
	// 		if len(archiveCommand) == 3 {
	// 			daysRangeInt, err = strconv.Atoi(archiveCommand[2])
	// 		}
	// 		if err != nil {
	// 			returnString = "Invalid Days Parameter"
	// 		} else {
	// 			go archiveChannelMedia(m, daysRangeInt)
	// 		}
	// 		sendMessage = true
	// 	}
	// }

	// Memes
	// if m.Author.ID == memeNameChangerUserID {
	// 	userNames := []string{
	// 		"Dumbbell Chrome Remover",
	// 		"Jay2Win",
	// 		"Perry",
	// 		"Frank Reynolds",
	// 		"Scraninator",
	// 		"Lord Scranian",
	// 		"Eddy Hall",
	// 		"Scran Master",
	// 		"Scran2D2",
	// 		"Bruce Scranner",
	// 		"Scranuel Jackson",
	// 		"Protein Baggins",
	// 		"Scran Solo",
	// 		"Jason Gainham",
	// 		"Captain Ameri-scran",
	// 		"Whoopi Swoleberg",
	// 		"Scranakin Skywalker",
	// 		"Obi-Wan Scranobi",
	// 		"The Swole Ranger",
	// 		"Gains Bond",
	// 		"Scranny G",
	// 		"Scranny Devito",
	// 	}
	// 	randomIndex := rand.Intn(len(userNames))
	// 	randomString := userNames[randomIndex]

	// 	err := dg.GuildMemberNickname(discordGuildID, memeNameChangerUserID, randomString)
	// 	if err != nil {
	// 		logger.Error().Err(err).Msg("Error changing nickname")
	// 		return
	// 	}
	// }

	// 8 Ball
	// if webhookEnabled && strings.HasPrefix(m.Content, commandPrefix+"8ball") {
	// 	// Trim the command part from the beginning
	// 	question := strings.TrimSpace(m.Content[len(commandPrefix+"8ball"):])
	// 	if question == "" {
	// 		// Handle case where no question was asked
	// 		returnString = "You need to ask a question!"
	// 		sendMessage = true
	// 	} else {
	// 		returnString = fmt.Sprintf("You asked: %s", question)
	// 		sendMessage = true
	// 	}

	// 	// Do something with the question
	// 	payload := WebhookPayload{
	// 		Question: question,
	// 		Message:  m,
	// 	}

	// 	data, err := json.Marshal(payload)
	// 	if err != nil {
	// 		log.Println("Error marshalling payload:", err)
	// 		return
	// 	}

	// 	resp, err := http.Post(
	// 		workflowsEndpoint,
	// 		"application/json",
	// 		bytes.NewBuffer(data),
	// 	)
	// 	if err != nil {
	// 		log.Println("Error sending webhook:", err)
	// 		return
	// 	}
	// 	defer resp.Body.Close()

	// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
	// 		log.Println("Unexpected webhook response status:", resp.Status)
	// 	}

	// }

	// Return the Message
	// if sendMessage {
	// 	s.ChannelMessageSend(m.ChannelID, returnString)
	// }
}

// type WebhookPayload struct {
// 	Question string                   `json:"question"`
// 	Message  *discordgo.MessageCreate `json:"message"`
// }

// func ready(s *discordgo.Session, event *discordgo.Ready) {
// 	// Set the playing status.
// 	s.UpdateGameStatus(0, "Lan Organization")
// }
