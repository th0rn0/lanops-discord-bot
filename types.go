package main

type HandleNewParticipantInput struct {
	Username  string `json:"username"`
	GiftedBy  string `json:"gifted_by"`
	DiscordID string `json:"discord_id"`
}

type HandleCreateEventInput struct {
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	URL     string `json:"url"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Address string `json:"address"`
}

type QueueMessage struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}
