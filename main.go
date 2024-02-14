package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello go")

	router := mux.NewRouter()

	http.ListenAndServe(":3000", router)
}
