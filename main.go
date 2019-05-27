package sendpulse

var grantType string
var clientID string
var clientSecret string

func initialize(rClientID string, rClientSecret string) {
	grantType = "client_credentials"
	clientID = rClientID
	clientSecret = rClientSecret
}

