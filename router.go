package main

import (
	"encoding/json"
	"net/http"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

// MethodGet path /hello
func (s *APIserver) handleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

// login
func (s *APIserver) handlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	// tys = LoginPayload
	payload := new(LoginPayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Payload"))
		return
	}

	kpayload := &KLoginPayload{
		clientId:     "rest-golang-auth",
		username:     payload.Username,
		password:     payload.Password,
		grantType:    "password",
		clientSecret: "FtkExkM8gvd5I5vtoK21U8fWbxtazGjd",
	}

	kres, err := s.client.login(kpayload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res := &LoginRes{
		AccessToken: kres.AccessToken,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}
