package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func pollWebRequests() {
	r := gin.Default()

	authorized := r.Group("", gin.BasicAuth(gin.Accounts{
		os.Getenv("AUTH_USER"): os.Getenv("AUTH_PASS"),
	}))

	r.Use(cors.Default())

	authorized.POST("/participants/new", handleNewParticipant)
	authorized.POST("/participants/gifted", handleGiftedParticipant)
	authorized.POST("/participants/transferred", handleTransferredParticipant)
	authorized.POST("/participants/remove", handleRemoveParticipant)

	authorized.POST("/events/create", handleCreateEvent)

	authorized.POST("/message/channel", handleMessageChannel)

	r.Run(":9888")
}
