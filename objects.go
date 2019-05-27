package sendpulse

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ErrorCode   int	   `json:"error_code"`
}

type emailArray struct {
	Content email `json:"email"`
}

type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type email struct {
	HTML    string 		`json:"html"`
	Text    string    	`json:"text"`
	Subject string      `json:"subject"`
	From    recipient   `json:"from"`
	To      []recipient `json:"to"`
}

type sendEmailResponse struct {
	Result bool `json:"result"`
}