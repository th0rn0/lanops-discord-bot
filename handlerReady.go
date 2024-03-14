package main

import "github.com/bwmarrin/discordgo"

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "!airhorn")
}
