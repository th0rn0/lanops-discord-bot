package main

import (
	"time"
)

func pollMessageQueue() {
	ticker := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			var queueMessages []QueueMessage
			db.Find(&queueMessages)
			for _, queueMessage := range queueMessages {
				_, err := dg.ChannelMessageSend(
					queueMessage.ChannelID,
					queueMessage.Message)
				if err != nil {
					logger.Error().Err(err).Msg("Error Sending Discord message")
					return
				}
				db.Delete(&queueMessage)
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
