package eightball

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	var returnString string
	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - 8Ball - Triggered", Level: "INFO"}

	question := strings.Join(args[:], " ")
	if question == "" {
		returnString = "You need to ask a question!"
	} else {
		returnString = fmt.Sprintf("You asked: %s", question)
		payload := WebhookPayload{
			Question: question,
			Message:  m,
		}

		data, err := json.Marshal(payload)
		if err != nil {
			msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
			returnString = "There was a error connecting to the API"
			return
		} else {
			resp, err := http.Post(
				cfg.EightBallEndpoint,
				"application/json",
				bytes.NewBuffer(data),
			)
			if err != nil {
				msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
				returnString = "There was a error connecting to the API"
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
				msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
				returnString = "There was a error connecting to the API"
			}
		}
	}

	s.ChannelMessageSend(m.ChannelID, returnString)
}

type WebhookPayload struct {
	Question string                   `json:"question"`
	Message  *discordgo.MessageCreate `json:"message"`
}
