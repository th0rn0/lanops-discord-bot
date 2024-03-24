package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleTransferredParticipant(c *gin.Context) {
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
