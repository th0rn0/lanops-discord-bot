package bot

import (
	"lanops/discord-bot/internal/bot/handlers"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"time"

	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func New(cfg config.Config, discordClient *discordgo.Session, db *gorm.DB, msgCh chan<- channels.MsgCh) (*Client, error) {
	client := &Client{
		cfg:   cfg,
		dg:    discordClient,
		db:    db,
		msgCh: msgCh,
	}

	// Register the Events for Discord Go
	client.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.OnMessage(s, m, cfg, msgCh)
	})
	client.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.OnReady(s, m, cfg, msgCh)
	})

	// Set the Intents
	client.dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	return client, nil
}

func (client *Client) Run() error {
	// Open the websocket
	if err := client.dg.Open(); err != nil {
		return err
	}
	defer client.dg.Close()

	go client.pollMessageQueue()

	// Wait for CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	return nil
}

type QueueMessage struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}

func (client *Client) pollMessageQueue() {
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			var queueMessages []QueueMessage
			client.db.Table("queue_messages").Find(&queueMessages)
			if len(queueMessages) != 0 {
				for _, queueMessage := range queueMessages {
					_, err := client.dg.ChannelMessageSend(
						queueMessage.ChannelID,
						queueMessage.Message)
					if err != nil {
						client.msgCh <- channels.MsgCh{Err: err, Message: "Error Sending Discord message", Level: "ERROR"}
						return
					}
					client.db.Delete(&queueMessage)
				}
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
