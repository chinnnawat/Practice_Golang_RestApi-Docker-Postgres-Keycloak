package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIserver struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}

func (s *APIserver) Run() {
	// s.migrate()
	router := mux.NewRouter()

	router.HandleFunc("/hello", s.handlerGreet)

	fmt.Printf("Server starting on Address : %s", s.addr)
	http.ListenAndServe(s.addr, router)
}

// func (s *APIserver) migrate() {
// 	// Migrate the schema
// 	s.db.AutoMigrate(&Post{})
// }
