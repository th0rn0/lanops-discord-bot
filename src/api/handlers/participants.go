package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func (s Client) NewParticipant(c *gin.Context) {
	var handleNewParticipantInput HandleNewParticipantInput
	if err := c.ShouldBindJSON(&handleNewParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"New Participant: %s", handleNewParticipantInput.Username))
	if !handleNewParticipantInput.NoMessage {
		addMessageToQueue(
			handleNewParticipantInput.ChannelID,
			fmt.Sprintf(
				"New Attendee: %s", handleNewParticipantInput.Username))
	}
	if handleNewParticipantInput.DiscordID != "" {
		if err := dg.GuildMemberRoleAdd(discordGuildID, handleNewParticipantInput.DiscordID, handleNewParticipantInput.RoleID); err != nil {
			logger.Error().Err(err).Msg("Error Updating User Permissions")
			c.JSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func (s Client) GiftedParticipant(c *gin.Context) {
	var handleGiftedParticipantInput HandleGiftedParticipantInput
	if err := c.ShouldBindJSON(&handleGiftedParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"Gifted Participant: %s", handleGiftedParticipantInput.Username))
	addMessageToQueue(
		handleGiftedParticipantInput.ChannelID,
		fmt.Sprintf(
			"New Attendee: %s - Gifted by %s - Such a Rinsing Geezer!", handleGiftedParticipantInput.Username, handleGiftedParticipantInput.GiftedBy))
	if handleGiftedParticipantInput.DiscordID != "" {
		if err := dg.GuildMemberRoleAdd(discordGuildID, handleGiftedParticipantInput.DiscordID, handleGiftedParticipantInput.RoleID); err != nil {
			logger.Error().Err(err).Msg("Error Updating User Permissions")
			c.JSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func (s Client) TransferredParticipant(c *gin.Context) {
	var handleTransferredParticipantInput HandleTransferredParticipantInput
	if err := c.ShouldBindJSON(&handleTransferredParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"Transferred Participant: %s", handleTransferredParticipantInput.Username))
	if handleTransferredParticipantInput.DiscordID != "" {
		if err := dg.GuildMemberRoleRemove(discordGuildID, handleTransferredParticipantInput.DiscordID, handleTransferredParticipantInput.RoleID); err != nil {
			logger.Error().Err(err).Msg("Error Updating User Permissions")
			c.JSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func (s Client) RemoveParticipant(c *gin.Context) {
	var handleRemoveParticipant HandleRemoveParticipant
	if err := c.ShouldBindJSON(&handleRemoveParticipant); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"Remove Participant: %s", handleRemoveParticipant.Username))
	if handleRemoveParticipant.DiscordID != "" {
		if err := dg.GuildMemberRoleRemove(discordGuildID, handleRemoveParticipant.DiscordID, handleRemoveParticipant.RoleID); err != nil {
			logger.Error().Err(err).Msg("Error Updating User Permissions")
			c.JSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, "OK")
}
