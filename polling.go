package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// func pollEventParticipantsApi(dg *discordgo.Session) {
// 	ticker := time.NewTicker(5 * time.Second)
// 	quit := make(chan struct{})

// 	db, _ := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})

// 	for {
// 		select {
// 		case <-ticker.C:
// 			apiEventParticipants := api.GetNextEventParticipants(apiUrl)
// 			// DEBUG
// 			for _, apiEventParticipant := range apiEventParticipants {

// 				var dbEventParticipant api.EventParticipant

// 				// Check if Participant exists, if it exists update if needed, otherwise add to DB
// 				result := db.First(&dbEventParticipant, apiEventParticipant.ID)

// 				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 					// We have a new Participant
// 					if err := db.Create(&api.EventParticipant{
// 						ID:       apiEventParticipant.ID,
// 						Username: apiEventParticipant.Username,
// 						Seat:     apiEventParticipant.Seat,
// 					}).Error; err != nil {
// 						// DEBUG - HANDLE THE ERRORS
// 						fmt.Println(err)
// 					}
// 					dg.ChannelMessageSend(
// 						channelID,
// 						fmt.Sprintf(
// 							"We have a new Participant! %s in seat %s",
// 							apiEventParticipant.Username,
// 							apiEventParticipant.Seat))
// 				} else {
// 					// fmt.Println("USER EXISTS")
// 					if dbEventParticipant.ID != apiEventParticipant.ID ||
// 						dbEventParticipant.Seat != apiEventParticipant.Seat ||
// 						dbEventParticipant.Username != apiEventParticipant.Username {
// 						// fmt.Println("SOMETHING HAS CHANGED")

// 						if err := db.Save(&api.EventParticipant{
// 							ID:       apiEventParticipant.ID,
// 							Username: apiEventParticipant.Username,
// 							Seat:     apiEventParticipant.Seat,
// 						}).Error; err != nil {
// 							// DEBUG - HANDLE THE ERRORS
// 							fmt.Println(err)
// 						}
// 						if dbEventParticipant.Seat != apiEventParticipant.Seat {
// 							dg.ChannelMessageSend(
// 								channelID,
// 								fmt.Sprintf(
// 									"%s is now in seat %s",
// 									apiEventParticipant.Username,
// 									apiEventParticipant.Seat))
// 						}
// 					}
// 				}

// 			}
// 			// DEBUG
// 			// fmt.Printf("body: %s", apiEventParticipants)
// 			// dg.ChannelMessageSend(channelID, fmt.Sprintf("%s", apiEventParticipants))
// 		case <-quit:
// 			ticker.Stop()
// 			return
// 		}
// 	}
// }

func pollMessageQueue(dg *discordgo.Session) {
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			var queueMessages []QueueMessage
			db.Find(&queueMessages)
			for _, queueMessage := range queueMessages {

				_, err := dg.ChannelMessageSend(
					channelID,
					queueMessage.Message)
				if err != nil {
					return
				}
				db.Delete(&queueMessage)
			}
			// log.Fatalln(results.Error)
			// fmt.Println(results)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
