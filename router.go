package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func handlerGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIserver) handlerGetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &Post{
		ID:        1,
		Title:     "golang",
		Content:   "awesome",
		CreatedAt: time.Now(),
	}
	json.NewEncoder(w).Encode(res)
}
