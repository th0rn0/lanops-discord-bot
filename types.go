package main

// INPUTS
type HandleNewParticipantInput struct {
	Username  string `json:"username"`
	DiscordID string `json:"discord_id"`
	ChannelID string `json:"channel_id"`
	RoleID    string `json:"role_id"`
	NoMessage bool   `json:"no_message"`
}

type HandleGiftedParticipantInput struct {
	HandleNewParticipantInput
	GiftedBy string `json:"gifted_by"`
}

type HandleTransferredParticipantInput struct {
	HandleNewParticipantInput
}

type HandleRemoveParticipant struct {
	HandleNewParticipantInput
}

type HandleCreateEventInput struct {
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	URL     string `json:"url"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Address string `json:"address"`
}

// OUTPUTS
type HandleCreateEventOutput struct {
	RoleID    string `json:"role_id"`
	ChannelID string `json:"channel_id"`
	EventID   string `json:"event_id"`
}

// MISC
type QueueMessage struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}
