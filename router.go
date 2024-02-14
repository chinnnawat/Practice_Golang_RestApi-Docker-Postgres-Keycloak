package main

import "net/http"

func handlerGreet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
