package dates

import (
	"fmt"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/manager"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	lanopsAPI manager.API
)

func formatDate(raw string) string {
	t, err := time.Parse("2006-01-02 15:04:05", raw)
	if err != nil {
		return raw
	}
	return t.Format("Mon 02 Jan 2006")
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Dates - Triggered", Level: "INFO"}

	lanopsAPI = manager.New(cfg.LanopsApiAddr)

	events, err := lanopsAPI.GetUpcomingEvents()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
		s.ChannelMessageSend(m.ChannelID, "There was an error connecting to the API")
		return
	}

	if len(events) == 0 {
		s.ChannelMessageSend(m.ChannelID, "No upcoming events found")
		return
	}

	msg := "**Upcoming Events**\n"
	for _, event := range events {
		msg += fmt.Sprintf("`%s` to `%s` — **%s**\n", formatDate(event.Start), formatDate(event.End), event.Name)
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}
