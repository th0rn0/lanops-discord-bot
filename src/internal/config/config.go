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

	discordAdminRoleId := os.Getenv("DISCORD_ADMIN_ROLE_ID")
	if discordAdminRoleId == "" {
		log.Fatal("❌ DISCORD_ADMIN_ROLE_ID not set in environment")
	}

	discordCommandPrefix := os.Getenv("DISCORD_COMMAND_PREFIX")
	if discordCommandPrefix == "" {
		log.Fatal("❌ DISCORD_COMMAND_PREFIX not set in environment")
	}

	discordGuildId := os.Getenv("DISCORD_GUILD_ID")
	if discordGuildId == "" {
		log.Fatal("❌ DISCORD_GUILD_ID not set in environment")
	}

	discordMemeNameChangerUserID := os.Getenv("MEME_NAME_CHANGER_USER_ID")
	if discordMemeNameChangerUserID == "" {
		log.Fatal("❌ MEME_NAME_CHANGER_USER_ID not set in environment")
	}

	discordArchiveChannelMediaRoleID := os.Getenv("ARCHIVE_CHANNEL_MEDIA_ROLE_ID")
	if discordArchiveChannelMediaRoleID == "" {
		log.Fatal("❌ ARCHIVE_CHANNEL_MEDIA_ROLE_ID not set in environment")
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
	lanOpsAPIAddr := os.Getenv("8BALL_ENDPOINT")
	if lanOpsAPIAddr == "" {
		log.Fatal("❌ 8BALL_ENDPOINT not set in environment")
	}

	discord := Discord{
		CommandPrefix:             discordCommandPrefix,
		Token:                     discordToken,
		GuildId:                   discordGuildId,
		AdminRoleId:               discordAdminRoleId,
		MemeNameChangerUserID:     discordMemeNameChangerUserID,
		ArchiveChannelMediaRoleID: discordArchiveChannelMediaRoleID,
	}

	api := ApiConfig{
		AdminUsername: apiAdminUsername,
		AdminPassword: apiAdminPassword,
		Port:          apiPort,
	}

	return Config{
		DbPath:            dbPath,
		Api:               api,
		LanopsAPIAddr:     lanOpsAPIAddr,
		EightBallEndPoint: eightBallEndPoint,
		Discord:           discord,
	}
}
