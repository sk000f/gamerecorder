package server

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer is a handler function for the gamerecorder
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Fred" {
		return "10"
	}

	return ""
}
