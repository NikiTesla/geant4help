package routes

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, let's check out what geant4 is"))
}

func help(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heeeeeeeeeeeeeeelp!"))
}
