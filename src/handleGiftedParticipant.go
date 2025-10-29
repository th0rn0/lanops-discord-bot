package main

// func handleGiftedParticipant(c *gin.Context) {
// 	var handleGiftedParticipantInput HandleGiftedParticipantInput
// 	if err := c.ShouldBindJSON(&handleGiftedParticipantInput); err != nil {
// 		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
// 		return
// 	}
// 	logger.Info().Msg(fmt.Sprintf(
// 		"Gifted Participant: %s", handleGiftedParticipantInput.Username))
// 	addMessageToQueue(
// 		handleGiftedParticipantInput.ChannelID,
// 		fmt.Sprintf(
// 			"New Attendee: %s - Gifted by %s - Such a Rinsing Geezer!", handleGiftedParticipantInput.Username, handleGiftedParticipantInput.GiftedBy))
// 	if handleGiftedParticipantInput.DiscordID != "" {
// 		if err := dg.GuildMemberRoleAdd(discordGuildID, handleGiftedParticipantInput.DiscordID, handleGiftedParticipantInput.RoleID); err != nil {
// 			logger.Error().Err(err).Msg("Error Updating User Permissions")
// 			c.JSON(http.StatusBadRequest, err)
// 		}
// 	}
// 	c.JSON(http.StatusOK, "OK")
// }
