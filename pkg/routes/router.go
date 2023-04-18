package routes

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/help", help)

	return rtr
}
