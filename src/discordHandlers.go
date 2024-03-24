package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// func connect(s *discordgo.Session, event *discordgo.Connect) {
// 	s.ChannelMessageSend(discordMainChannelID, "I AM AWAKE AT LAST")
// }

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

	// Return the Message
	if sendMessage {
		s.ChannelMessageSend(m.ChannelID, returnString)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "BOT GO BRRR")
}
