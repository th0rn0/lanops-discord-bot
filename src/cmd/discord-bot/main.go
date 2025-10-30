package main

import (
	"fmt"
	"lanops/discord-bot/api"
	"lanops/discord-bot/internal/bot"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/msgqueue"

	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	logger zerolog.Logger
	cfg    config.Config
	msgCh  = make(chan channels.MsgCh, 20)
	db     *gorm.DB
)

func main() {
	logger = zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
	logger.Info().Msg("Initializing LanOps Discord Bot")

	logger.Info().Msg("Loading Config")
	cfg = config.Load()

	logger.Info().Msg("Loading Database")
	db, err := gorm.Open(sqlite.Open(cfg.DbPath), &gorm.Config{})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed Connecting to DB")
	}
	db.AutoMigrate(&msgqueue.Message{})

	go func() {
		for msg := range msgCh {
			if msg.Err != nil {
				logger.Error().Err(msg.Err).Msg(msg.Message)
			} else {
				logger.Info().Msg(msg.Message)
			}
		}
	}()

	logger.Info().Msg("Starting Message Queue")
	msgQueue := msgqueue.New(db, cfg)

	logger.Info().Msg("Starting Discord Session")
	discordSession, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed start Discord Session")
	}

	logger.Info().Msg("Starting API")
	api := api.SetupRouter(msgQueue, discordSession, cfg, logger)
	go func() {
		api.Run(fmt.Sprintf(":%s", cfg.Api.Port))
	}()

	logger.Info().Msg("Starting LanOps Discord Bot")

	botClient, err := bot.New(cfg, discordSession, msgQueue, msgCh)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create bot")
	}

	if err := botClient.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start bot")
	}

}
