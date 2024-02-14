package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	addr string
}

func NewAPIServer(addr string) {
	return &APIserver{
		addr: addr,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handlerGreet)

	fmt.Println("Server starting on Address : %s", s.addr)
	http.ListenAndServe(s.addr, router)
}
