package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lanops/discord-bot/api"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	commandPrefix string = "!"
	channelID     string = "411444570266796043"
)

var (
	token          string
	apiUrl         string
	dg             *discordgo.Session
	db             *gorm.DB
	discordGuildID string
)

func main() {
	// DEBUG - put this into init()
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token = os.Getenv("DISCORD_TOKEN")
	apiUrl = os.Getenv("API_URL")
	discordGuildID = os.Getenv("DISCORD_SERVER_ID")

	// Load DB
	db, err = gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&api.EventParticipant{})
	db.AutoMigrate(&QueueMessage{})

	// Create a new Discord session using the provided bot token.
	dg, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Register the Events for Discord Go
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(connect)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Start Listeners and Polling
	go startWebRouter()
	// go pollEventParticipantsApi(dg)
	go pollMessageQueue(dg)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("LanOps Discord Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func startWebRouter() {
	r := gin.Default()

	r.Use(cors.Default())

	// DEBUG - split out into new/gifted/tranferred/refunded etc
	r.POST("/webhooks/participants", handleNewParticipant)

	r.POST("/webhooks/events/create", handleCreateEvent)

	r.Run(":8888")
}

func handleNewParticipant(c *gin.Context) {
	var handleNewParticipantInput HandleNewParticipantInput
	if err := c.ShouldBindJSON(&handleNewParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	if handleNewParticipantInput.GiftedBy != "" {
		addMessageToQueue(
			channelID,
			fmt.Sprintf(
				"New attendee: %s - Gifted by %s", handleNewParticipantInput.Username, handleNewParticipantInput.GiftedBy))
	} else {
		addMessageToQueue(
			channelID,
			fmt.Sprintf(
				"New attendee: %s", handleNewParticipantInput.Username))
	}
	if err := dg.GuildMemberRoleAdd(discordGuildID, handleNewParticipantInput.DiscordID, "1217587354592739438"); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, "OK")
}

func addMessageToQueue(channelID string, message string) {
	dbMessage := QueueMessage{
		ChannelID: channelID,
		Message:   message,
	}
	// DEBUG - DO SOME ERROR HANDLING
	db.Where(QueueMessage{ChannelID: channelID, Message: message}).FirstOrCreate(&dbMessage)
}

// Function to create new event, a new role and save them to a database?
func handleCreateEvent(c *gin.Context) {
	var handleCreateEventInput HandleCreateEventInput
	if err := c.ShouldBindJSON(&handleCreateEventInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	// Format Dates
	var startTime, _ = time.Parse("2006-01-02 15:04:00", handleCreateEventInput.Start)
	var endTime, _ = time.Parse("2006-01-02 15:04:00", handleCreateEventInput.End)

	discordRole, err := dg.GuildRoleCreate(discordGuildID, &discordgo.RoleParams{
		Name: handleCreateEventInput.Name + " Participant",
	})
	if err != nil {
		fmt.Println(err)
	}

	discordChannel, err := dg.GuildChannelCreate(discordGuildID, handleCreateEventInput.Slug, 0)
	if err != nil {
		fmt.Println(err)
	}

	discordEvent, err := dg.GuildScheduledEventCreate(discordGuildID, &discordgo.GuildScheduledEventParams{
		Name:               handleCreateEventInput.Name,
		Description:        handleCreateEventInput.URL,
		ScheduledStartTime: &startTime,
		ScheduledEndTime:   &endTime,
		EntityType:         3,
		PrivacyLevel:       2,
		EntityMetadata: &discordgo.GuildScheduledEventEntityMetadata{
			Location: handleCreateEventInput.Address,
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = dg.ChannelMessage(discordChannel.ID, "@everyone first - all you're event are belong to us!")
	if err != nil {
		fmt.Println(err)
	}

	handleCreateEventOutput := HandleCreateEventOutput{
		RoleID:    discordRole.ID,
		ChannelID: discordChannel.ID,
		EventID:   discordEvent.ID,
	}
	c.JSON(http.StatusOK, handleCreateEventOutput)
}

type HandleCreateEventOutput struct {
	RoleID    string `json:"role_id"`
	ChannelID string `json:"channel_id"`
	EventID   string `json:"event_id"`
}
