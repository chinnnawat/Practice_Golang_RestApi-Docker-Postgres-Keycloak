package main

import "net/http"

type KeycloakService interface {
	login(payload *KLoginPayload) error
}

type KLoginPayload struct {
	clientId     string
	username     string
	password     string
	grantType    string
	clientSecret string
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client struct {
	httpClient *http.Client
}

type KLoginRes struct {
	AccessToken string `json:"access_token"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
}
