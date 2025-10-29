package bot

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type Client struct {
	cfg config.Config
	dg  *discordgo.Session
	db  *gorm.DB
	// streams streams.Client
	msgCh chan<- channels.MsgCh
}
