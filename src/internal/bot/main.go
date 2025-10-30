package bot

import (
	"lanops/discord-bot/internal/bot/handlers"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/internal/msgqueue"
	"time"

	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func New(cfg config.Config, discordSession *discordgo.Session, msgQueue msgqueue.Client, msgCh chan<- channels.MsgCh) (*Client, error) {
	client := &Client{
		cfg:      cfg,
		dg:       discordSession,
		msgQueue: msgQueue,
		msgCh:    msgCh,
	}

	// Register the Events for Discord Go
	client.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.OnMessage(s, m, cfg, msgCh)
	})
	client.dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.OnReady(s)
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

	go client.pollMsgQueue()

	// Wait for CTRL+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	return nil
}

func (client *Client) pollMsgQueue() {
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			msgQueue := client.msgQueue.Get()
			if len(msgQueue) != 0 {
				for _, queueMessage := range msgQueue {
					_, err := client.dg.ChannelMessageSend(
						queueMessage.ChannelID,
						queueMessage.Content)
					if err != nil {
						client.msgCh <- channels.MsgCh{Err: err, Message: "Error Sending Discord message", Level: "ERROR"}
						return
					}
					client.msgQueue.Delete(queueMessage)
				}
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
