package main

import (
	"lanops/discord-bot/api"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	commandPrefix string = "!"
)

var (
	token                string
	lanopsAPI            api.API
	dg                   *discordgo.Session
	db                   *gorm.DB
	discordGuildID       string
	discordMainChannelID string
	logger               zerolog.Logger
)

func init() {
	var err error

	logger = zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
	logger.Info().Msg("Initializing LanOps Discord Bot")

	// Env Variables
	logger.Info().Msg("Loading Environment Variables")
	godotenv.Load()
	lanopsAPI = api.New(os.Getenv("API_URL"))
	token = os.Getenv("DISCORD_TOKEN")
	discordGuildID = os.Getenv("DISCORD_SERVER_ID")
	discordMainChannelID = os.Getenv("DISCORD_MAIN_CHANNEL_ID")

	// Database
	logger.Info().Msg("Connecting to Database")
	db, err = gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	if err != nil {
		logger.Fatal().Err(err).Msg("Error Connecting to Database")
	}
	db.AutoMigrate(&QueueMessage{})

	// Discord Bot
	logger.Info().Msg("Connecting to Discord API")
	dg, err = discordgo.New("Bot " + token)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error Creating Discord Session")
	}
	logger.Info().Msg("Initalization Complete")
}
func main() {
	logger.Info().Msg("Starting Discord Bot")

	// Register the Events for Discord Go
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	// dg.AddHandler(connect)

	// Set the Intents
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket
	err := dg.Open()
	if err != nil {
		logger.Fatal().Err(err).Msg("Error Opening Discord session")
	}

	// Start Listeners and Polling
	logger.Info().Msg("Starting GIN Web Server")
	go pollWebRequests()
	logger.Info().Msg("Polling for Messages...")
	go pollMessageQueue()

	logger.Info().Msg("LanOps Discord Bot is now running.  Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
