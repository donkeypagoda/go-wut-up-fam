package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /api", api)

	log.Print("starting server....")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
