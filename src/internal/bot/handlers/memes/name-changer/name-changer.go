package start

import (
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"lanops/discord-bot/internal/jukebox"
	"slices"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	// if slices.Contains(m.Member.Roles, cfg.Discord.AdminRoleId) {
	// 	var returnString string
	// 	msgCh <- channels.MsgCh{Err: nil, Message: "Message Event - Jukebox Start - Triggered", Level: "INFO"}
	// 	jukeboxClient := jukebox.New(cfg)
	// 	if err := jukeboxClient.Start(); err != nil {
	// 		msgCh <- channels.MsgCh{Err: err, Message: "Something went wrong", Level: "ERROR"}
	// 		returnString = "There was a error connecting to the API"
	// 	} else {
	// 		returnString = "Starting Jukebox"
	// 	}
	// 	s.ChannelMessageSend(m.ChannelID, returnString)
	// }

		// TODO - this wont go in the message create part? maybe?
	if m.Author.ID == memeNameChangerUserID {
		userNames := []string{
			"Dumbbell Chrome Remover",
			"Jay2Win",
			"Perry",
			"Frank Reynolds",
			"Scraninator",
			"Lord Scranian",
			"Eddy Hall",
			"Scran Master",
			"Scran2D2",
			"Bruce Scranner",
			"Scranuel Jackson",
			"Protein Baggins",
			"Scran Solo",
			"Jason Gainham",
			"Captain Ameri-scran",
			"Whoopi Swoleberg",
			"Scranakin Skywalker",
			"Obi-Wan Scranobi",
			"The Swole Ranger",
			"Gains Bond",
			"Scranny G",
			"Scranny Devito",
		}
		randomIndex := rand.Intn(len(userNames))
		randomString := userNames[randomIndex]

		err := dg.GuildMemberNickname(discordGuildID, memeNameChangerUserID, randomString)
		if err != nil {
			logger.Error().Err(err).Msg("Error changing nickname")
			return
		}
}
