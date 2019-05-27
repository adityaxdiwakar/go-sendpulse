package sendpulse

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"bytes"
)

var grantType string
var clientID string
var clientSecret string

func initialize(rClientID string, rClientSecret string) {
	grantType = "client_credentials"
	clientID = rClientID
	clientSecret = rClientSecret
}

func getKey() {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type": grantType,
		"client_id": clientID,
		"client_secret": clientSecret,
	})
	
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://api.sendpulse.com/oauth/access_token", 
						   "application/json",
						   bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}