package main

import (
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// var returnString string

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "getevents" {
		// var upcomingEvent APIResponseEventUpcoming
		// resp, err := http.Get(apiUrl + "/events/upcoming")
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// fmt.Printf("body: %s", body)
		// // fmt.Print(body)
		// // var arr []string
		// err = json.Unmarshal(body, &upcomingEvent)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// fmt.Println(upcomingEvent)
		// s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("body: %s", body))
	}

	if m.Content == "getUpcoming" {

	}
}
