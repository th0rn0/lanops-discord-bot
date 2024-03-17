package main

func addMessageToQueue(channelID string, message string) {
	dbMessage := QueueMessage{
		ChannelID: channelID,
		Message:   message,
	}
	// DEBUG - DO SOME ERROR HANDLING
	db.Where(QueueMessage{ChannelID: channelID, Message: message}).FirstOrCreate(&dbMessage)
}
