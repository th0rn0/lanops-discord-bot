package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type HandleCreateEventInput struct {
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	URL     string `json:"url"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Address string `json:"address"`
}

type HandleCreateEventOutput struct {
	RoleID    string `json:"role_id"`
	ChannelID string `json:"channel_id"`
	EventID   string `json:"event_id"`
}

func (s Client) CreateEvent(c *gin.Context) {
	var handleCreateEventInput HandleCreateEventInput
	if err := c.ShouldBindJSON(&handleCreateEventInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}

	s.logger.Info().Msg(fmt.Sprintf("New Event: %s", handleCreateEventInput.Name))

	var startTime, _ = time.Parse("2006-01-02 15:04:00", handleCreateEventInput.Start)
	var endTime, _ = time.Parse("2006-01-02 15:04:00", handleCreateEventInput.End)

	discordRole, err := s.discordSession.GuildRoleCreate(s.cfg.Discord.GuildId, &discordgo.RoleParams{
		Name: handleCreateEventInput.Name + " Participant",
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Error Creating Guild Role")
	}

	discordChannel, err := s.discordSession.GuildChannelCreate(s.cfg.Discord.GuildId, handleCreateEventInput.Slug, 0)
	if err != nil {
		s.logger.Error().Err(err).Msg("Error Creating Guild Channel")
	}

	discordEvent, err := s.discordSession.GuildScheduledEventCreate(s.cfg.Discord.GuildId, &discordgo.GuildScheduledEventParams{
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
		s.logger.Error().Err(err).Msg("Error Creating Guild Scheduled Event")
	}

	_, err = s.discordSession.ChannelMessageSend(discordChannel.ID, "First - all your event are belong to us!")
	if err != nil {
		s.logger.Error().Err(err).Msg("Error Sending Discord message")
	}

	handleCreateEventOutput := HandleCreateEventOutput{
		RoleID:    discordRole.ID,
		ChannelID: discordChannel.ID,
		EventID:   discordEvent.ID,
	}

	c.JSON(http.StatusOK, handleCreateEventOutput)
}
