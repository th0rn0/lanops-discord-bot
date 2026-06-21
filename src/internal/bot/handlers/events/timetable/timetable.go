package timetable

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/manager"

	"github.com/bwmarrin/discordgo"
)

var (
	lanopsAPI manager.API
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	var returnString string
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Timetable - Triggered", Level: "INFO"}

	lanopsAPI = manager.New(cfg.LanopsApiAddr)

	var event, err = lanopsAPI.GetNextEvent()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
		returnString = "There was a error connecting to the API"
	} else {
		returnString = "\nTimetable for Next Event:\n"
		returnString += event.Name + "\n"
		returnString += event.URL.Base + event.URL.Timetables
	}
	s.ChannelMessageSend(m.ChannelID, returnString)
}
