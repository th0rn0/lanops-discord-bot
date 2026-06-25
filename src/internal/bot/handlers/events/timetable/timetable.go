package timetable

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

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Timetable - Triggered", Level: "INFO"}

	lanopsAPI = manager.New(cfg.LanopsApiAddr)

	event, err := lanopsAPI.GetNextEvent()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
		s.ChannelMessageSend(m.ChannelID, "There was an error connecting to the API")
		return
	}

	timetables, err := lanopsAPI.GetEventTimetables(event.Slug)
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong fetching timetables", Level: "ERROR"}
		s.ChannelMessageSend(m.ChannelID, "There was an error fetching the timetable")
		return
	}

	if len(timetables) == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("No timetables published yet for %s", event.Name))
		return
	}

	msg := fmt.Sprintf("**Timetable for %s**\n", event.Name)
	for _, timetable := range timetables {
		if len(timetables) > 1 {
			msg += fmt.Sprintf("\n__%s__\n", timetable.Name)
		}
		for _, entry := range timetable.Data {
			t, parseErr := time.Parse("2006-01-02 15:04:05", entry.StartTime)
			var timeStr string
			if parseErr != nil {
				timeStr = entry.StartTime
			} else {
				timeStr = t.Format("Mon 02 Jan 15:04")
			}
			msg += fmt.Sprintf("`%s` **%s** — %s\n", timeStr, entry.Name, entry.Description)
		}
	}

	s.ChannelMessageSend(m.ChannelID, msg)
}
