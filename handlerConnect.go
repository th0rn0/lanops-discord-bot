package main

import "github.com/bwmarrin/discordgo"

func connect(s *discordgo.Session, event *discordgo.Connect) {
	s.ChannelMessageSend(channelID, "I AM AWAKE AT LAST")
}
