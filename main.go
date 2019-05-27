package sendpulse

import (
	"encoding/json"
	"encoding/base64"
	"net/http"
	"io/ioutil"
	"bytes"
	"errors"
)

var grantType string
var clientID string
var clientSecret string
var accessToken string
var fromContact recipient

func initialize(rClientID string, rClientSecret string, rName string, rFromEmail string) {
	grantType = "client_credentials"
	clientID = rClientID
	clientSecret = rClientSecret
	fromContact = recipient{
		Name: rName,
		Email: rFromEmail,
	}
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

func SendEmail(html []byte, text []byte, subject string, to []recipient) error {
	encoded := base64.StdEncoding.EncodeToString(html)
	mailObj := emailArray{
		email{
			HTML: encoded,
			Text: string(text),
			Subject: subject,
			From: fromContact,
			To: to,
		},
	}
	
	mailByteSlice, err := json.Marshal(mailObj)
	if err != nil {
		return errors.New("Something wrong with email object -> string")
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.sendpulse.com/smtp/emails", 
		bytes.NewBuffer([]byte(mailByteSlice)),
	)
	
	if err != nil {
		return errors.New("Something wrong with string -> request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("Something wrong sending the email")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var emailResponse sendEmailResponse
	err = json.Unmarshal([]byte(body), &emailResponse)
	if err != nil {
		return errors.New("Repopulation issue")
	}

	if !emailResponse.Result {
		return errors.New("Something went wrong at SendPulse")
	}
	
	return nil
}