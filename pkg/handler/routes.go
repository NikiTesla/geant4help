package handler

import (
	"net/http"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, let's check out what geant4 is"))
}

func (h *Handler) help(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heeeeeeeeeeeeeeelp!"))
}
