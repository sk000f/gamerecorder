package main

import (
	"log"
	"net/http"

	"github.com/sk000f/gamerecorder/pkg/server"
)

func main() {
	server := &server.PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

// InMemoryPlayerStore is a simple in memory store of player scores
type InMemoryPlayerStore struct{}

// GetPlayerScore returns the score for the specified player from the InMemoryPlayerStore
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
