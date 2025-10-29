package dates

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	var returnString string
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Dates - Triggered", Level: "INFO"}

	var nextEvent, err = lanopsAPI.GetUpcomingEvents()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
	} else {
		returnString = formatUpcomingEventDatesMessage(nextEvent)
	}
	s.ChannelMessageSend(m.ChannelID, returnString)
}
