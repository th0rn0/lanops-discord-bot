package bot

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/msgqueue"

	"github.com/bwmarrin/discordgo"
)

type Client struct {
	cfg      config.Config
	dg       *discordgo.Session
	msgQueue msgqueue.Client
	msgCh    chan<- channels.MsgCh
}
