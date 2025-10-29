package main

// func handleRemoveParticipant(c *gin.Context) {
// 	var handleRemoveParticipant HandleRemoveParticipant
// 	if err := c.ShouldBindJSON(&handleRemoveParticipant); err != nil {
// 		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
// 		return
// 	}
// 	logger.Info().Msg(fmt.Sprintf(
// 		"Remove Participant: %s", handleRemoveParticipant.Username))
// 	if handleRemoveParticipant.DiscordID != "" {
// 		if err := dg.GuildMemberRoleRemove(discordGuildID, handleRemoveParticipant.DiscordID, handleRemoveParticipant.RoleID); err != nil {
// 			logger.Error().Err(err).Msg("Error Updating User Permissions")
// 			c.JSON(http.StatusBadRequest, err)
// 		}
// 	}
// 	c.JSON(http.StatusOK, "OK")
// }
