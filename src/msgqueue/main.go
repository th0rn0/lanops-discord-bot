package msgqueue

import (
	"lanops/discord-bot/internal/config"

	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg config.Config) Client {
	return Client{
		db:  db,
		cfg: cfg,
	}
}

func (c Client) Create(channelID string, message string) {
	dbMessage := Message{
		ChannelID: channelID,
		Content:   message,
	}
	c.db.Where(Message{ChannelID: channelID, Content: message}).FirstOrCreate(&dbMessage)
}

func (c Client) Get() (messages []Message) {
	c.db.Table("messages").Find(&messages)
	return messages
}

func (c Client) Delete(message Message) {
	c.db.Delete(&message)
}
