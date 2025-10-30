package attendees

import (
	"fmt"
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
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Attendees - Triggered", Level: "INFO"}

	lanopsAPI = manager.New(cfg.LanopsApiAddr)

	var participants, err = lanopsAPI.GetNextEventParticipants()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
		returnString = "There was a error connecting to the API"
	} else {
		returnString = "\nAttendees for Next Event:\n"
		for _, participant := range participants {
			returnString += fmt.Sprintf(
				"%s - %s\n", participant.Username, participant.Seat)
		}
	}
	s.ChannelMessageSend(m.ChannelID, returnString)
}
