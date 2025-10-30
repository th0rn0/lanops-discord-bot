package handlers

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"

	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/rand"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate, cfg config.Config, msgCh chan<- channels.MsgCh) {
	// On all Messages
	if m.Author.Bot {
		return
	}

	if m.Author.ID == cfg.Discord.MemeNameChangerUserId {
		userNames := []string{
			"Dumbbell Chrome Remover",
			"Jay2Win",
			"Perry",
			"Frank Reynolds",
			"Scraninator",
			"Lord Scranian",
			"Eddy Hall",
			"Scran Master",
			"Scran2D2",
			"Bruce Scranner",
			"Scranuel Jackson",
			"Protein Baggins",
			"Scran Solo",
			"Jason Gainham",
			"Captain Ameri-scran",
			"Whoopi Swoleberg",
			"Scranakin Skywalker",
			"Obi-Wan Scranobi",
			"The Swole Ranger",
			"Gains Bond",
			"Scranny G",
			"Scranny Devito",
		}
		randomIndex := rand.Intn(len(userNames))
		randomString := userNames[randomIndex]
		err := s.GuildMemberNickname(cfg.Discord.GuildId, cfg.Discord.MemeNameChangerUserId, randomString)
		if err != nil {
			msgCh <- channels.MsgCh{Err: err, Message: "Error changing nickname", Level: "ERROR"}
		}
	}

	if !strings.HasPrefix(m.Content, cfg.Discord.CommandPrefix) {
		return
	}

	content := strings.TrimPrefix(m.Content, cfg.Discord.CommandPrefix)
	parts := strings.Fields(content)
	if len(parts) == 0 {
		return
	}

	// On Command Messages
	for i := len(parts); i > 0; i-- {
		key := strings.Join(parts[:i], " ")
		if handler, ok := Registry[key]; ok {
			handler(s, m, parts[:i], parts[i:], cfg, msgCh)
			return
		}
	}
}

func OnReady(s *discordgo.Session) {
	// Set the playing status.
	s.UpdateGameStatus(0, "Lan Partying!")
}
