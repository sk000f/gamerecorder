package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer is the gamerecorder server
type PlayerServer struct {
	Store PlayerStore
}

// PlayerStore holds player score data
type PlayerStore interface {
	GetPlayerScore(name string) int
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.Store.GetPlayerScore(player))
}

// GetPlayerScore returns the specified player's score
func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Fred" {
		return "10"
	}

	return ""
}
