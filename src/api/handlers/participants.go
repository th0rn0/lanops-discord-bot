package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (s Client) NewParticipant(c *gin.Context) {
	var handleNewParticipantInput HandleNewParticipantInput
	if err := c.ShouldBindJSON(&handleNewParticipantInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	s.logger.Info().Msg(fmt.Sprintf(
		"New Participant: %s", handleNewParticipantInput.Username))
	if !handleNewParticipantInput.NoMessage {
		s.msgQueue.Create(
			handleNewParticipantInput.ChannelID,
			fmt.Sprintf(
				"New Attendee: %s", handleNewParticipantInput.Username))
	}
	if handleNewParticipantInput.DiscordID != "" {
		if err := s.discordSession.GuildMemberRoleAdd(s.cfg.Discord.GuildId, handleNewParticipantInput.DiscordID, handleNewParticipantInput.RoleID); err != nil {
			s.logger.Error().Err(err).Msg("Error Updating User Permissions")
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
	s.logger.Info().Msg(fmt.Sprintf(
		"Gifted Participant: %s", handleGiftedParticipantInput.Username))
	s.msgQueue.Create(
		handleGiftedParticipantInput.ChannelID,
		fmt.Sprintf(
			"New Attendee: %s - Gifted by %s - Such a Rinsing Geezer!", handleGiftedParticipantInput.Username, handleGiftedParticipantInput.GiftedBy))
	if handleGiftedParticipantInput.DiscordID != "" {
		if err := s.discordSession.GuildMemberRoleAdd(s.cfg.Discord.GuildId, handleGiftedParticipantInput.DiscordID, handleGiftedParticipantInput.RoleID); err != nil {
			s.logger.Error().Err(err).Msg("Error Updating User Permissions")
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
	s.logger.Info().Msg(fmt.Sprintf(
		"Transferred Participant: %s", handleTransferredParticipantInput.Username))
	if handleTransferredParticipantInput.DiscordID != "" {
		if err := s.discordSession.GuildMemberRoleRemove(s.cfg.Discord.GuildId, handleTransferredParticipantInput.DiscordID, handleTransferredParticipantInput.RoleID); err != nil {
			s.logger.Error().Err(err).Msg("Error Updating User Permissions")
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
	s.logger.Info().Msg(fmt.Sprintf(
		"Remove Participant: %s", handleRemoveParticipant.Username))
	if handleRemoveParticipant.DiscordID != "" {
		if err := s.discordSession.GuildMemberRoleRemove(s.cfg.Discord.GuildId, handleRemoveParticipant.DiscordID, handleRemoveParticipant.RoleID); err != nil {
			s.logger.Error().Err(err).Msg("Error Updating User Permissions")
			c.JSON(http.StatusBadRequest, err)
		}
	}
	c.JSON(http.StatusOK, "OK")
}
