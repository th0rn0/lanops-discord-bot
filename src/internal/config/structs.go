package config

type Config struct {
	DbPath            string
	Api               ApiConfig
	LanopsApiAddr     string
	EightBallEndpoint string
	Discord           Discord
	MediaArchiver     MediaArchiver
}

type ApiConfig struct {
	AdminUsername   string
	AdminPassword   string
	Port            string
	AuthCallBackUrl string
}

type Discord struct {
	CommandPrefix         string
	Token                 string
	GuildId               string
	MemeNameChangerUserId string
	MediaArchiverRoleId   string
}

type MediaArchiver struct {
	GoogleDriveUploadDirId string
	TmpUploadDirPath       string
	GoogleDriveKey         string
}
