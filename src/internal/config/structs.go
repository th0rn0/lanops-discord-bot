package config

type Config struct {
	DbPath            string
	Api               ApiConfig
	LanopsAPIAddr     string
	EightBallEndPoint string
	Discord           Discord
}

type ApiConfig struct {
	AdminUsername   string
	AdminPassword   string
	Port            string
	AuthCallBackUrl string
}

type Discord struct {
	CommandPrefix             string
	Token                     string
	GuildId                   string
	AdminRoleId               string
	MemeNameChangerUserID     string
	ArchiveChannelMediaRoleID string
}
