package msgqueue

import (
	"lanops/discord-bot/internal/config"

	"gorm.io/gorm"
)

type Client struct {
	db  *gorm.DB
	cfg config.Config
}

type Message struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ChannelID string `json:"channel_id"`
	Content   string `json:"content"`
}
