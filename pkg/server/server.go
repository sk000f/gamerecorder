package server

import (
	"fmt"
	"net/http"
)

// PlayerServer is a handler function for the gamerecorder
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
