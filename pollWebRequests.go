package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func pollWebRequests() {
	r := gin.Default()

	r.Use(cors.Default())

	// DEBUG - split out into new/gifted/tranferred/refunded etc
	r.POST("/webhooks/participants", handleNewParticipant)

	r.POST("/webhooks/events/create", handleCreateEvent)

	r.Run(":8888")
}
