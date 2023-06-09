package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) InitRouter() *mux.Router {
	// authorization and registration
	rtr := mux.NewRouter()

	auth := rtr.NewRoute().Subrouter()
	auth.HandleFunc("/", h.indexPage)
	auth.HandleFunc("/help", h.helpPage)
	auth.HandleFunc("/login", h.logInPage).Methods("GET")
	auth.HandleFunc("/signup", h.signUpPage).Methods("GET")
	auth.HandleFunc("/login", h.logIn).Methods("POST")
	auth.HandleFunc("/signup", h.signUp).Methods("POST")
	auth.HandleFunc("/signout", h.signOut).Methods("GET")

	urtr := rtr.PathPrefix("/user").Subrouter()
	urtr.Use(h.authMiddleware)
	urtr.HandleFunc("", h.userIndexPage)
	urtr.HandleFunc("/", h.userIndexPage)
	urtr.HandleFunc("/profile", h.userPage)
	urtr.HandleFunc("/edit_information", h.editUserInfoPage).Methods("GET")
	urtr.HandleFunc("/edit_information", h.editUserInfo).Methods("POST")
	urtr.HandleFunc("/codegen", h.codeGenPage).Methods("GET")
	urtr.HandleFunc("/codegen", h.codeGen).Methods("POST")

	rtr.PathPrefix("/web/static").Handler(
		http.StripPrefix("/web/static", http.FileServer(http.Dir("./web/static/"))))

	return rtr
}
