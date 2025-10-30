package next

import (
	"fmt"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/manager"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var (
	lanopsAPI manager.API
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	var returnString string
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Get Events - Triggered", Level: "INFO"}

	lanopsAPI = manager.New(cfg.LanopsApiAddr)

	var event, err = lanopsAPI.GetNextEvent()
	if err != nil {
		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
		returnString = "There was a error connecting to the API"
	} else {
		returnString = "\nNext Event:\n"
		returnString += event.Name + "\n"
		returnString += event.Start + " to " + event.End + "\n"
		returnString += event.Description.Short + "\n"
		returnString += fmt.Sprintf(
			"Seats: %s/%s \n",
			strconv.Itoa(len(event.Participants)),
			strconv.Itoa(event.Capacity))
		returnString += "BOOK NOW!: " + event.URL.Base + event.URL.Tickets
	}
	s.ChannelMessageSend(m.ChannelID, returnString)
}
