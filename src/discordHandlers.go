package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
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
