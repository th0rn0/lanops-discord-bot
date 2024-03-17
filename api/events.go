package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func (a API) GetNextEvent() (Event, error) {
	return getNextEvent(a.URL)
}

func getNextEvent(url string) (Event, error) {
	var nextEvent Event
	resp, err := http.Get(url + "/events/next")
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Get API Response")
		return nextEvent, err
	}
	if resp.StatusCode != 200 {
		errorMessage := fmt.Sprintf(
			"Status Code From %s: %s",
			url,
			strconv.Itoa(resp.StatusCode))
		logger.Error().Msg(errorMessage)
		return nextEvent, errors.New(errorMessage)
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &nextEvent)
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Marshal JSON from Response")
		return nextEvent, err
	}
	return nextEvent, nil
}

func (a API) GetNextEventParticipants() ([]EventParticipant, error) {
	var participants []EventParticipant
	var nextEvent, err = getNextEvent(a.URL)
	if err != nil {
		return participants, err
	}
	participants = nextEvent.Participants
	return participants, nil
}

func (a API) GetUpcomingEvents() ([]Event, error) {
	var upcomingEvents []Event
	resp, err := http.Get(a.URL + "/events/upcoming")
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Get API Response")
		return upcomingEvents, err
	}
	if resp.StatusCode != 200 {
		errorMessage := fmt.Sprintf(
			"Status Code From %s: %s",
			a.URL,
			strconv.Itoa(resp.StatusCode))
		logger.Error().Msg(errorMessage)
		return upcomingEvents, errors.New(errorMessage)
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &upcomingEvents)
	if err != nil {
		logger.Error().Err(err).Msg("Cannot Marshal JSON from Response")
		return upcomingEvents, err
	}
	return upcomingEvents, nil
}
