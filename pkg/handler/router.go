package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) InitRouter() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", h.index)
	rtr.HandleFunc("/help", h.help)

	rtr.PathPrefix("/web/static").Handler(
		http.StripPrefix("/web/static", http.FileServer(http.Dir("./web/static/"))))

	return rtr
}
