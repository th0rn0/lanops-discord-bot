package main

func addMessageToQueue(channelID string, message string) {
	dbMessage := QueueMessage{
		ChannelID: channelID,
		Message:   message,
	}
	db.Where(QueueMessage{ChannelID: channelID, Message: message}).FirstOrCreate(&dbMessage)
}
