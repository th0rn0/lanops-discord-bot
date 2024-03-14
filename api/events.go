package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetNextEventParticipants(apiUrl string) APIResponseEventParticipants {
	var participants APIResponseEventParticipants
	resp, err := http.Get(apiUrl + "/events/asdasd/participants")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("body: %s", body)
	err = json.Unmarshal(body, &participants)
	if err != nil {
		log.Fatalln(err)
	}
	return participants
}
