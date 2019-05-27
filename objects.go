package main

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ErrorCode   int	   `json:"error_code"`
}