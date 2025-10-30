package handlers

import (
	"lanops/discord-bot/internal/bot/handlers/eightball"
	eventsAttendees "lanops/discord-bot/internal/bot/handlers/events/attendees"
	eventsDates "lanops/discord-bot/internal/bot/handlers/events/dates"
	eventsNext "lanops/discord-bot/internal/bot/handlers/events/next"
	"lanops/discord-bot/internal/bot/handlers/help"
	"lanops/discord-bot/internal/bot/handlers/mediaarchiver"

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
	Register("event next", eventsNext.Handler)
	Register("event dates", eventsDates.Handler)
	Register("event attendees", eventsAttendees.Handler)
	// Media Archiver
	Register("media archive", mediaarchiver.Handler)
	// 8Ball
	Register("8ball", eightball.Handler)
}
