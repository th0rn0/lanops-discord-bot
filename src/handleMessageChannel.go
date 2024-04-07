package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleMessageChannel(c *gin.Context) {
	var handleMessageChannelInput HandleMessageChannelInput
	if err := c.ShouldBindJSON(&handleMessageChannelInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	logger.Info().Msg(fmt.Sprintf(
		"Send Channel Message: %s", handleMessageChannelInput.Message))
	addMessageToQueue(
		handleMessageChannelInput.ChannelID,
		handleMessageChannelInput.Message)
	c.JSON(http.StatusOK, "OK")
}
