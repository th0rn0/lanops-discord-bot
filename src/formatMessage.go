package main

import (
	"fmt"
	"lanops/discord-bot/manager"
	"strconv"
)

func formatEventParticipantsMessage(participants []manager.EventParticipant) string {
	var returnMessage = "\nAttendees for Next Event:\n"
	for _, participant := range participants {
		returnMessage += fmt.Sprintf(
			"%s - %s\n", participant.Username, participant.Seat)
	}
	return returnMessage
}

func formatUpcomingEventDatesMessage(events []manager.Event) string {
	var returnMessage = "\nUpcoming Events:\n"
	for _, event := range events {
		returnMessage += fmt.Sprintf(
			"%s - %s to %s\n", event.Name, event.Start, event.End)
	}
	return returnMessage
}

func formatNextEventMessage(event manager.Event) string {
	var returnMessage = "\nNext Event:\n"
	returnMessage += event.Name + "\n"
	returnMessage += event.Start + " to " + event.End + "\n"
	returnMessage += event.Description.Short + "\n"
	returnMessage += fmt.Sprintf(
		"Seats: %s/%s \n",
		strconv.Itoa(len(event.Participants)),
		strconv.Itoa(event.Capacity))
	returnMessage += "BOOK NOW!: " + event.URL.Base + event.URL.Tickets
	return returnMessage
}
