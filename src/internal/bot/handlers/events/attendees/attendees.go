package attendees

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	var returnString string
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Attendees - Triggered", Level: "INFO"}

	var nextEvent, err = lanopsAPI.GetNextEventParticipants()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
	} else {
		returnString = formatEventParticipantsMessage(nextEvent)
	}
	s.ChannelMessageSend(m.ChannelID, returnString)
}
