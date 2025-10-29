package handlers

import (
	eventsAttendees "lanops/discord-bot/internal/bot/handlers/events/attendees"
	eventsDates "lanops/discord-bot/internal/bot/handlers/events/dates"
	eventsGet "lanops/discord-bot/internal/bot/handlers/events/get"
	help "lanops/discord-bot/internal/bot/handlers/help"

	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"

	"github.com/bwmarrin/discordgo"
)

type HandlerFunc func(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	commandParts []string,
	args []string,
	cfg config.Config,
	msgCh chan<- channels.MsgCh,
)

var Registry = map[string]HandlerFunc{}

func Register(command string, handler HandlerFunc) {
	Registry[command] = handler
}

func init() {
	// Help
	Register("help", help.Handler)
	// Events
	Register("get event", eventsGet.Handler)
	Register("get dates", eventsDates.Handler)
	Register("get attendees", eventsAttendees.Handler)
	// Media Archiver
	// TODO - Add me for the rest
	Register("get attendees", eventsAttendees.Handler)
	// Misc
	Register("get attendees", eventsAttendees.Handler)
	Register("get attendees", eventsAttendees.Handler)

}
