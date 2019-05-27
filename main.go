package sendpulse

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
	"errors"
)

var grantType string
var clientID string
var clientSecret string
var accessToken string


type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ErrorCode   int	   `json:"error_code"`
}

func initialize(rClientID string, rClientSecret string) {
	grantType = "client_credentials"
	clientID = rClientID
	clientSecret = rClientSecret
}

func getKey() (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type": grantType,
		"client_id": clientID,
		"client_secret": clientSecret,
	})
	
	if err != nil {
		return "", errors.New("Marshalling request payload gave an error.")
	}

	resp, err := http.Post("https://api.sendpulse.com/oauth/access_token", 
						   "application/json",
						   bytes.NewBuffer(requestBody))

	if err != nil {
		return "", errors.New("Making the request gave an error.")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Reading the response gave an error.")
	}

	var response oauthTokenResponse
	err = json.Unmarshal([]byte(body), &response)

	if response.ErrorCode != 0 {
		return "", errors.New("SendPulse sent an error.")
	}

	accessToken = response.AccessToken
	return response.AccessToken, nil
}

