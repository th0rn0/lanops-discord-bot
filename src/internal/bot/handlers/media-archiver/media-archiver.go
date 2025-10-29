package start

import (
	"lanops/party-discord-bot/internal/channels"
	"lanops/party-discord-bot/internal/config"
	"slices"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm/logger"
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

	if slices.Contains(m.Member.Roles, archiveChannelMediaRoleID) {
		if strings.HasPrefix(m.Content, commandPrefix+"media archive") {
			logger.Info().Msg("Message Create Event - Image Archive - Triggered")
			returnString = "Archiving Channel Media!"
			var err error
			daysRangeInt := 0
			archiveCommand := strings.Split(m.Content, " ")
			if len(archiveCommand) == 3 {
				daysRangeInt, err = strconv.Atoi(archiveCommand[2])
			}
			if err != nil {
				returnString = "Invalid Days Parameter"
			} else {
				go archiveChannelMedia(m, daysRangeInt)
			}
			sendMessage = true
		}
	}
}
