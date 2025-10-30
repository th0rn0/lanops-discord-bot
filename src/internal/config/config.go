package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() Config {
	godotenv.Load()

	// DB
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("❌ DB_PATH not set in environment")
	}

	// Discord
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("❌ DISCORD_TOKEN not set in environment")
	}

	discordCommandPrefix := os.Getenv("DISCORD_COMMAND_PREFIX")
	if discordCommandPrefix == "" {
		log.Fatal("❌ DISCORD_COMMAND_PREFIX not set in environment")
	}

	discordGuildId := os.Getenv("DISCORD_GUILD_ID")
	if discordGuildId == "" {
		log.Fatal("❌ DISCORD_GUILD_ID not set in environment")
	}

	discordMemeNameChangerUserID := os.Getenv("DISCORD_MEME_NAME_CHANGER_USER_ID")
	if discordMemeNameChangerUserID == "" {
		log.Fatal("❌ DISCORD_MEME_NAME_CHANGER_USER_ID not set in environment")
	}

	discordMediaArchiverRoleID := os.Getenv("DISCORD_MEDIA_ARCHIVER_ROLE_ID")
	if discordMediaArchiverRoleID == "" {
		log.Fatal("❌ DISCORD_MEDIA_ARCHIVER_ROLE_ID not set in environment")
	}

	// API
	apiAdminUsername := os.Getenv("API_ADMIN_USERNAME")
	if apiAdminUsername == "" {
		log.Fatal("❌ API_ADMIN_USERNAME not set in environment")
	}
	apiAdminPassword := os.Getenv("API_ADMIN_PASSWORD")
	if apiAdminPassword == "" {
		log.Fatal("❌ API_ADMIN_PASSWORD not set in environment")
	}
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		log.Fatal("❌ API_PORT not set in environment")
	}

	// EightBall Endpoint
	eightBallEndPoint := os.Getenv("8BALL_ENDPOINT")
	if eightBallEndPoint == "" {
		log.Fatal("❌ 8BALL_ENDPOINT not set in environment")
	}

	// LanOps API Address
	lanOpsAPIAddr := os.Getenv("LANOPS_API_ADDR")
	if lanOpsAPIAddr == "" {
		log.Fatal("❌ LANOPS_API_ADDR not set in environment")
	}

	// Media Archiver
	mediaArchiverGoogleDriveUploadDirId := os.Getenv("MEDIA_ARCHIVER_GOOGLE_DRIVE_UPLOAD_DIR_ID")
	if mediaArchiverGoogleDriveUploadDirId == "" {
		log.Fatal("❌ MEDIA_ARCHIVER_GOOGLE_DRIVE_UPLOAD_DIR_ID not set in environment")
	}

	mediaArchiverTempUploadDirPath := os.Getenv("MEDIA_ARCHIVER_TEMP_UPLOAD_DIR_PATH")
	if mediaArchiverTempUploadDirPath == "" {
		log.Fatal("❌ MEDIA_ARCHIVER_TEMP_UPLOAD_DIR_PATH not set in environment")
	}

	mediaArchiverGoogleDriveKey := os.Getenv("MEDIA_ARCHIVER_GOOGLE_DRIVE_KEY_JSON_BASE64")
	if mediaArchiverGoogleDriveKey == "" {
		log.Fatal("❌ MEDIA_ARCHIVER_GOOGLE_DRIVE_KEY_JSON_BASE64 not set in environment")
	}

	discord := Discord{
		CommandPrefix:         discordCommandPrefix,
		Token:                 discordToken,
		GuildId:               discordGuildId,
		MemeNameChangerUserId: discordMemeNameChangerUserID,
		MediaArchiverRoleId:   discordMediaArchiverRoleID,
	}

	api := ApiConfig{
		AdminUsername: apiAdminUsername,
		AdminPassword: apiAdminPassword,
		Port:          apiPort,
	}

	mediaArchiver := MediaArchiver{
		GoogleDriveUploadDirId: mediaArchiverGoogleDriveUploadDirId,
		TmpUploadDirPath:       mediaArchiverTempUploadDirPath,
		GoogleDriveKey:         mediaArchiverGoogleDriveKey,
	}

	return Config{
		DbPath:            dbPath,
		Api:               api,
		LanopsApiAddr:     lanOpsAPIAddr,
		EightBallEndpoint: eightBallEndPoint,
		Discord:           discord,
		MediaArchiver:     mediaArchiver,
	}
}
