package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func connect(s *discordgo.Session, event *discordgo.Connect) {
	s.ChannelMessageSend(discordMainChannelID, "I AM AWAKE AT LAST")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// var returnString string

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == commandPrefix+"get event" {
		logger.Info().Msg("Message Create Event - Get Event - Triggered")
		// DEBUG - Error Handling
		var nextEvent = lanopsAPI.GetNextEvent()
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("body: %s", nextEvent))
	}

	if m.Content == commandPrefix+"get dates" {
		logger.Info().Msg("Message Create Event - Get Dates - Triggered")
		var upcomingEvents = lanopsAPI.GetUpcomingEvents()
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("body: %s", upcomingEvents))
	}

	if m.Content == commandPrefix+"get attendees" {
		logger.Info().Msg("Message Create Event - Get Attendees - Triggered")
		var participants = lanopsAPI.GetNextEventParticipants()
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("body: %s", participants))
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "BOT GO BRRR")
}
