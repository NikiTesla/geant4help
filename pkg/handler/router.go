package handler

import (
	"github.com/gorilla/mux"
)

func (h Handler) InitRouter() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", h.index)
	rtr.HandleFunc("/help", h.help)

	return rtr
}
