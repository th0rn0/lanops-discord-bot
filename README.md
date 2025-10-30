# lanops-discord-bot

Discord bot for LanOps - Written in GO.

## Prerequisites

- ```cp src/.env.example src/.env``` and fill it in

#### Install Dependencies

```bash
cd src
go mod tidy
```

## Usage

Entry Point:
```bash
go run ./cmd/discord-bot
```

## Bot Commands

| Command           | Input           | Description                  |
|-------------------|-----------------|------------------------------|
| `event next`      | None            | Lists all available streams. |
| `event dates`     | None            | Enables a specific stream.   |
| `event attendees` | None            | Disables a specific stream.  |
| `media archive`   | Days to Archive | Starts the jukebox playback. |
| `8ball`           | Question        | Stops the jukebox playback.  |

## Env

| Variable                                      | Description                                          |
|-----------------------------------------------|------------------------------------------------------|
| `DB_PATH`                                     | Path to DB.                                          |
| `DISCORD_TOKEN`                               | Token for authenticating the Discord bot.            |
| `DISCORD_COMMAND_PREFIX`                      | Prefix used for bot commands in Discord (e.g., `!`). |
| `DISCORD_GUILD_ID`                            | ID of the GUILD.                                     |
| `DISCORD_MEME_NAME_CHANGER_USER_ID`           | ID of the User to Randomize Names.                   |
| `DISCORD_MEDIA_ARCHIVER_ROLE_ID`              | ID of the MEDIA ARCHIVER Role on Discord.            |
| `API_ADMIN_USERNAME`                          | Username for authenticating with the BOT API.        |
| `API_ADMIN_PASSWORD`                          | Password for authenticating with the BOT API.        |
| `API_PORT`                                    | Port of the BOT API.                                 |
| `8BALL_ENDPOINT`                              | Base URL for the 8Ball.                              |
| `LANOPS_API_ADDR`                             | Base URL for the LANOPS Manager API.                 |
| `MEDIA_ARCHIVER_GOOGLE_DRIVE_UPLOAD_DIR_ID`   | Google Drive Directory ID to upload archived media.  |
| `MEDIA_ARCHIVER_TEMP_UPLOAD_DIR_PATH`         | Temp directory for archived media.                   |
| `MEDIA_ARCHIVER_GOOGLE_DRIVE_KEY_JSON_BASE64` | Base64 Encoded Authentication Key for Google Drive.  |

## Docker

```docker build -f resources/docker/Dockerfile .```

```
docker run -d \
  --name discord-bot \
  --restart unless-stopped \
  -e DB_PATH= \
  -e DISCORD_TOKEN= \
  -e DISCORD_COMMAND_PREFIX= \
  -e DISCORD_GUILD_ID= \
  -e DISCORD_MEME_NAME_CHANGER_USER_ID= \
  -e DISCORD_MEDIA_ARCHIVER_ROLE_ID= \
  -e API_ADMIN_USERNAME= \
  -e API_ADMIN_PASSWORD= \
  -e API_PORT= \
  -e 8BALL_ENDPOINT= \
  -e LANOPS_API_ADDR= \
  -e MEDIA_ARCHIVER_GOOGLE_DRIVE_UPLOAD_DIR_ID= \
  -e MEDIA_ARCHIVER_TEMP_UPLOAD_DIR_PATH= \
  -e MEDIA_ARCHIVER_GOOGLE_DRIVE_KEY_JSON_BASE64= \
  -p 8888:8888 \
  -v /mnt/servdata/lanops/jukebox/db:/db \
  th0rn0/lanops-discord-bot:latest
```

```
  discord-bot:
    image: th0rn0/lanops-discord-bot:latest
    container_name: discord-bot
    restart: unless-stopped
    environment:
      DB_PATH: 
      DISCORD_TOKEN: 
      DISCORD_COMMAND_PREFIX: 
      DISCORD_GUILD_ID: 
      DISCORD_MEME_NAME_CHANGER_USER_ID: 
      DISCORD_MEDIA_ARCHIVER_ROLE_ID: 
      API_ADMIN_USERNAME: 
      API_ADMIN_PASSWORD: 
      API_PORT: 
      8BALL_ENDPOINT: 
      LANOPS_API_ADDR: 
      MEDIA_ARCHIVER_GOOGLE_DRIVE_UPLOAD_DIR_ID: 
      MEDIA_ARCHIVER_TEMP_UPLOAD_DIR_PATH: 
      MEDIA_ARCHIVER_GOOGLE_DRIVE_KEY_JSON_BASE64: 
```