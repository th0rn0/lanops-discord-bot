package main

import (
	"fmt"
	"lanops/discord-bot/api"
	"strconv"

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
		s.ChannelMessageSend(m.ChannelID, formatNextEventMessage(nextEvent))
	}

	if m.Content == commandPrefix+"get dates" {
		logger.Info().Msg("Message Create Event - Get Dates - Triggered")
		var upcomingEvents = lanopsAPI.GetUpcomingEvents()
		s.ChannelMessageSend(m.ChannelID, formatUpcomingEventDatesMessage(upcomingEvents))
	}

	if m.Content == commandPrefix+"get attendees" {
		logger.Info().Msg("Message Create Event - Get Attendees - Triggered")
		var participants = lanopsAPI.GetNextEventParticipants()
		s.ChannelMessageSend(m.ChannelID, formatEventParticipantsMessage(participants))
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "BOT GO BRRR")
}

func formatEventParticipantsMessage(participants []api.EventParticipant) string {
	var returnMessage = "\nAttendees for Next Event:\n"
	for _, participant := range participants {
		returnMessage += fmt.Sprintf(
			"%s - %s\n", participant.Username, participant.Seat)
	}
	return returnMessage
}

func formatUpcomingEventDatesMessage(events []api.Event) string {
	var returnMessage = "\nUpcoming Events:\n"
	for _, event := range events {
		returnMessage += fmt.Sprintf(
			"%s - %s to %s\n", event.Name, event.Start, event.End)
	}
	return returnMessage
}

func formatNextEventMessage(event api.Event) string {
	var returnMessage = "\nNext Event:\n"
	returnMessage += event.Name + "\n"
	returnMessage += event.Start + " to " + event.End + "\n"
	returnMessage += event.Description.Short + "\n"
	returnMessage += fmt.Sprintf("Seats: %s \n", strconv.Itoa(event.Capacity))
	returnMessage += "BOOK NOW!: " + event.URL.Base + event.URL.Tickets
	return returnMessage
}
