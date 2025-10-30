package api

import (
	"os"
	"time"

	"lanops/discord-bot/api/handlers"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/msgqueue"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func SetupRouter(msgQueue msgqueue.Client, discordSession *discordgo.Session, cfg config.Config, logger zerolog.Logger) *gin.Engine {
	logger.Info().Msg("Loading API")
	gin.DefaultWriter = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Dur("latency", latency).
			Msg("request handled")
	})

	r.Use(cors.Default())

	// Handlers
	handlers := handlers.New(msgQueue, discordSession, cfg, logger)
	authorized := r.Group("", gin.BasicAuth(gin.Accounts{
		cfg.Api.AdminUsername: cfg.Api.AdminPassword,
	}))

	authorized.POST("/participants/new", handlers.NewParticipant)
	authorized.POST("/participants/gifted", handlers.GiftedParticipant)
	authorized.POST("/participants/transferred", handlers.TransferredParticipant)
	authorized.POST("/participants/remove", handlers.RemoveParticipant)

	authorized.POST("/events/create", handlers.CreateEvent)

	authorized.POST("/message/channel", handlers.MessageChannel)

	return r
}
