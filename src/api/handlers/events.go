package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func (s Client) CreateEvent(c *gin.Context) {
	var handleCreateEventInput HandleCreateEventInput
	if err := c.ShouldBindJSON(&handleCreateEventInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}

	logger.Info().Msg(fmt.Sprintf(
		"New Event: %s", handleCreateEventInput.Name))

	// Format Dates
	var startTime, _ = time.Parse("2006-01-02 15:04:00", handleCreateEventInput.Start)
	var endTime, _ = time.Parse("2006-01-02 15:04:00", handleCreateEventInput.End)

	discordRole, err := dg.GuildRoleCreate(discordGuildID, &discordgo.RoleParams{
		Name: handleCreateEventInput.Name + " Participant",
	})
	if err != nil {
		logger.Error().Err(err).Msg("Error Creating Guild Role")
	}

	discordChannel, err := dg.GuildChannelCreate(discordGuildID, handleCreateEventInput.Slug, 0)
	if err != nil {
		logger.Error().Err(err).Msg("Error Creating Guild Channel")
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
		logger.Error().Err(err).Msg("Error Creating Guild Scheduled Event")
	}

	_, err = dg.ChannelMessageSend(discordChannel.ID, "First - all your event are belong to us!")
	if err != nil {
		logger.Error().Err(err).Msg("Error Sending Discord message")
	}

	handleCreateEventOutput := HandleCreateEventOutput{
		RoleID:    discordRole.ID,
		ChannelID: discordChannel.ID,
		EventID:   discordEvent.ID,
	}

	c.JSON(http.StatusOK, handleCreateEventOutput)
}
