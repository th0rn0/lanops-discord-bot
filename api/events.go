package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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

func (a API) GetNextEvent() Event {
	return getNextEvent(a.URL)
}

func getNextEvent(url string) Event {
	var nextEvent Event
	resp, err := http.Get(url + "/events/next")
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Get API Response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &nextEvent)
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Marshal JSON from Response")
	}
	return nextEvent
}

func (a API) GetNextEventParticipants() []EventParticipant {
	var nextEvent = getNextEvent(a.URL)
	return nextEvent.Participants
}

func (a API) GetUpcomingEvents() []Event {
	var upcomingEvents []Event
	resp, err := http.Get(a.URL + "/events/upcoming")
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Get API Response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &upcomingEvents)
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Marshal JSON from Response")
	}
	return upcomingEvents
}
