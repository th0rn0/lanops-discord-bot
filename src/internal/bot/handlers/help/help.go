package help

import (
	"fmt"
	"lanops/discord-bot/internal/channels"
	"lanops/discord-bot/internal/config"
	"sort"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate, commandParts []string, args []string, cfg config.Config, msgCh chan<- channels.MsgCh) {
	helpMsg := "**Available Commands:**\n"
	helpMsg += formatCommands(map[string]string{
		cfg.Discord.CommandPrefix + "event next":      "Get Next Event",
		cfg.Discord.CommandPrefix + "event attendees": "Get Next Event Attendees",
		cfg.Discord.CommandPrefix + "event dates":     "Get all Event Dates",
		cfg.Discord.CommandPrefix + "8ball":           "Ask a 8Ball Question",
	})
	s.ChannelMessageSend(m.ChannelID, helpMsg)
}

func formatCommands(cmds map[string]string) string {
	keys := make([]string, 0, len(cmds))
	for k := range cmds {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var out string
	for _, k := range keys {
		out += fmt.Sprintf("`%s` - %s\n", k, cmds[k])
	}
	return out
}
