package main

import (
	"lanops/discord-bot/jukebox"
	"lanops/discord-bot/manager"
	"os"
	"os/signal"
	"strconv"
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
	token                        string
	lanopsAPI                    manager.API
	jukeboxAPI                   jukebox.API
	dg                           *discordgo.Session
	db                           *gorm.DB
	discordGuildID               string
	discordMainChannelID         string
	discordJukeBoxControlRoleID  string
	discordJukeboxControlEnabled bool
	logger                       zerolog.Logger
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
	lanopsAPI = manager.New(os.Getenv("MANAGER_URL"))
	jukeboxAPI = jukebox.New(os.Getenv("JUKEBOX_URL"), os.Getenv("JUKEBOX_USERNAME"), os.Getenv("JUKEBOX_PASSWORD"))
	token = os.Getenv("DISCORD_TOKEN")
	discordGuildID = os.Getenv("DISCORD_SERVER_ID")
	discordMainChannelID = os.Getenv("DISCORD_MAIN_CHANNEL_ID")
	discordJukeBoxControlRoleID = os.Getenv("DISCORD_JUKEBOX_CONTROL_ROLE_ID")
	discordJukeboxControlEnabled, _ = strconv.ParseBool(os.Getenv("DISCORD_JUKEBOX_CONTROL_ENABLED"))

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
