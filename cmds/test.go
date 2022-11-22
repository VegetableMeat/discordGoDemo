package cmds

import "github.com/bwmarrin/discordgo"

// Test testコマンド
func Test(s *discordgo.Session, m *discordgo.MessageCreate, array []string) {
	s.ChannelMessageSend(m.ChannelID, "test")
}
