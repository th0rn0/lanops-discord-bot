package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandleMessageChannelInput struct {
	ChannelID string `json:"channel_id"`
	Content   string `json:"message"`
}

func (s Client) MessageChannel(c *gin.Context) {
	var handleMessageChannelInput HandleMessageChannelInput
	if err := c.ShouldBindJSON(&handleMessageChannelInput); err != nil {
		c.JSON(http.StatusInternalServerError, "Cannot Marshal JSON")
		return
	}
	s.logger.Info().Msg(fmt.Sprintf(
		"Send Channel Message: %s", handleMessageChannelInput.Content))
	s.msgQueue.Create(
		handleMessageChannelInput.ChannelID,
		handleMessageChannelInput.Content)
	c.JSON(http.StatusOK, "OK")
}
