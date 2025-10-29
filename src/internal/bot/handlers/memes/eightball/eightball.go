package start

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	// if slices.Contains(m.Member.Roles, cfg.Discord.AdminRoleId) {
	// 	var returnString string
	// 	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Jukebox Start - Triggered", Level: "INFO"}
	// 	jukeboxClient := jukebox.New(cfg)
	// 	if err := jukeboxClient.Start(); err != nil {
	// 		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
	// 		returnString = "There was a error connecting to the API"
	// 	} else {
	// 		returnString = "Starting Jukebox"
	// 	}
	// 	s.ChannelMessageSend(m.ChannelID, returnString)
	// }

	// Trim the command part from the beginning
	question := strings.TrimSpace(m.Content[len(commandPrefix+"8ball"):])
	if question == "" {
		// Handle case where no question was asked
		returnString = "You need to ask a question!"
		sendMessage = true
	} else {
		returnString = fmt.Sprintf("You asked: %s", question)
		sendMessage = true
	}

	// Do something with the question
	payload := WebhookPayload{
		Question: question,
		Message:  m,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return
	}

	resp, err := http.Post(
		workflowsEndpoint,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.Println("Error sending webhook:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Println("Unexpected webhook response status:", resp.Status)
	}
}

type WebhookPayload struct {
	Question string                   `json:"question"`
	Message  *discordgo.MessageCreate `json:"message"`
}
