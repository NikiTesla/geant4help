package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) InitRouter() *mux.Router {
	// authorization and registration
	rtr := mux.NewRouter()

	auth := rtr.NewRoute().Subrouter()
	auth.HandleFunc("/", h.index)
	auth.HandleFunc("/login", h.LogInPage).Methods("GET")
	auth.HandleFunc("/signup", h.SignUpPage).Methods("GET")
	auth.HandleFunc("/login", h.logIn).Methods("POST")
	auth.HandleFunc("/signup", h.signUp).Methods("POST")

	urtr := rtr.PathPrefix("/user").Subrouter()
	urtr.Use(h.authMiddleware)
	urtr.HandleFunc("", h.help)
	urtr.HandleFunc("/", h.help)

	rtr.PathPrefix("/web/static").Handler(
		http.StripPrefix("/web/static", http.FileServer(http.Dir("./web/static/"))))

	return rtr
}
