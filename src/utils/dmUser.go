package utils

import (
	"github.com/bwmarrin/discordgo"
)

func DmUser(dg *discordgo.Session, userID string, message string) error {
	channel, err := dg.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	_, err = dg.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return err
	}
	return nil
}
