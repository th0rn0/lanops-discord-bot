package main

import (
	"time"
)

func pollMessageQueue() {
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			var queueMessages []QueueMessage
			db.Table("queue_messages").Find(&queueMessages)
			if len(queueMessages) != 0 {
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
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
