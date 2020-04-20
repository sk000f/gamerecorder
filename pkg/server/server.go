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
	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
		break
	case http.MethodGet:
		p.showScore(w, r)
		break
	}

}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
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
