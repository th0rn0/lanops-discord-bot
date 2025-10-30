package handlers

import (
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/internal/msgqueue"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

type Client struct {
	msgQueue       msgqueue.Client
	discordSession *discordgo.Session
	cfg            config.Config
	logger         zerolog.Logger
}

func New(msgQueue msgqueue.Client, discordSession *discordgo.Session, cfg config.Config, logger zerolog.Logger) Client {
	return Client{
		msgQueue:       msgQueue,
		discordSession: discordSession,
		cfg:            cfg,
		logger:         logger,
	}
}
