package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleNewParticipant(c *gin.Context) {
	var handleNewParticipantInput HandleNewParticipantInput
	if err := c.ShouldBindJSON(&handleNewParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"New Participant: %s", handleNewParticipantInput.Username))
	addMessageToQueue(
		handleNewParticipantInput.ChannelID,
		fmt.Sprintf(
			"New Attendee: %s", handleNewParticipantInput.Username))
	if err := dg.GuildMemberRoleAdd(discordGuildID, handleNewParticipantInput.DiscordID, handleNewParticipantInput.RoleID); err != nil {
		logger.Error().Err(err).Msg("Error Updating User Permissions")
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "OK")
}

func handleGiftedParticipant(c *gin.Context) {
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
	if err := dg.GuildMemberRoleAdd(discordGuildID, handleGiftedParticipantInput.DiscordID, handleGiftedParticipantInput.RoleID); err != nil {
		logger.Error().Err(err).Msg("Error Updating User Permissions")
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "OK")
}

func handleTransferredParticipant(c *gin.Context) {
	var handleTransferredParticipantInput HandleTransferredParticipantInput
	if err := c.ShouldBindJSON(&handleTransferredParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"Transferred Participant: %s", handleTransferredParticipantInput.Username))
	if err := dg.GuildMemberRoleRemove(discordGuildID, handleTransferredParticipantInput.DiscordID, handleTransferredParticipantInput.RoleID); err != nil {
		logger.Error().Err(err).Msg("Error Updating User Permissions")
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "OK")
}
