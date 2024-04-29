package jukebox

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
)

func init() {
	logger = zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()
}

func New(URL string) API {
	a := API{URL: URL}
	return a
}

func (a API) Control(command string) string {
	var returnString string
	client := &http.Client{}
	req, err := http.NewRequest("POST", a.URL+"/player/"+command, nil)
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Get API Response")
	}
	req.SetBasicAuth("admin", "changeme")
	switch command {
	case "start":
		logger.Info().Msg("Jukebox Control - START")
		resp, err := client.Do(req)
		if err != nil {
			logger.Error().Err(err).Msg("Cannot Get API Response")
		}
		if resp.StatusCode != 202 {
			errorMessage := fmt.Sprintf(
				"Status Code From %s: %s",
				a.URL+"/player/start",
				strconv.Itoa(resp.StatusCode))
			logger.Error().Msg(errorMessage)
			returnString = "Something went wrong " + errorMessage
		} else {
			returnString = "Jukebox - Started Polling Spotify"
		}
	case "stop":
		logger.Info().Msg("Jukebox Control - STOP")
		resp, err := client.Do(req)
		if err != nil {
			logger.Error().Err(err).Msg("Cannot Get API Response")
		}
		if resp.StatusCode != 202 {
			errorMessage := fmt.Sprintf(
				"Status Code From %s: %s",
				a.URL+"/player/stop",
				strconv.Itoa(resp.StatusCode))
			logger.Error().Msg(errorMessage)
			returnString = "Something went wrong " + errorMessage
		} else {
			returnString = "Jukebox - Stop Polling Spotify"
		}
	case "skip":
		logger.Info().Msg("Jukebox Control - SKIP")
		resp, err := client.Do(req)
		if err != nil {
			logger.Error().Err(err).Msg("Cannot Get API Response")
		}
		if resp.StatusCode != 202 {
			errorMessage := fmt.Sprintf(
				"Status Code From %s: %s",
				a.URL+"/player/skip",
				strconv.Itoa(resp.StatusCode))
			logger.Error().Msg(errorMessage)
			returnString = "Something went wrong " + errorMessage
		} else {
			returnString = "Jukebox - Song Skipped"
		}
	default:
		logger.Info().Msg("Jukebox Control - Unknown Command")
		returnString = "Command not Recognized"
	}

	return returnString
}
