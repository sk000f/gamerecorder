package main

import (
	"log"
	"net/http"

	"github.com/sk000f/gamerecorder/pkg/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
