package main

// func handleNewParticipant(c *gin.Context) {
// 	var handleNewParticipantInput HandleNewParticipantInput
// 	if err := c.ShouldBindJSON(&handleNewParticipantInput); err != nil {
// 		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
// 		return
// 	}
// 	logger.Info().Msg(fmt.Sprintf(
// 		"New Participant: %s", handleNewParticipantInput.Username))
// 	if !handleNewParticipantInput.NoMessage {
// 		addMessageToQueue(
// 			handleNewParticipantInput.ChannelID,
// 			fmt.Sprintf(
// 				"New Attendee: %s", handleNewParticipantInput.Username))
// 	}
// 	if handleNewParticipantInput.DiscordID != "" {
// 		if err := dg.GuildMemberRoleAdd(discordGuildID, handleNewParticipantInput.DiscordID, handleNewParticipantInput.RoleID); err != nil {
// 			logger.Error().Err(err).Msg("Error Updating User Permissions")
// 			c.JSON(http.StatusBadRequest, err)
// 		}
// 	}
// 	c.JSON(http.StatusOK, "OK")
// }
