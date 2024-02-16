package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	addr   string
	client Client
}

func NewAPIServer(addr string) *APIserver {
	httpClient := &http.Client{}
	return &APIserver{
		addr: addr,
		client: Client{
			httpClient: httpClient,
		},
	}
}

func (s *APIserver) Run() {
	// s.migrate()
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handleGreet)
	router.HandleFunc("/login", s.handlerLogin)

	fmt.Printf("Server starting on Address : %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
