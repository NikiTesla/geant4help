package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/NikiTesla/geant4help/pkg/repository"
)

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/index.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/footer.html", h.Env.Config.StaticDir),
	)
	if err != nil {
		h.Env.Logger.Error("can't parse template index.html")
	}

	tmpl.ExecuteTemplate(w, "index", nil)

}

func (h *Handler) help(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/help.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/footer.html", h.Env.Config.StaticDir),
	)
	if err != nil {
		h.Env.Logger.Error("can't parse template index.html")
	}

	tmpl.ExecuteTemplate(w, "help", nil)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Env.Logger.Error("cannot parse form value 'age'")
		return
	}
	salary, err := strconv.ParseFloat(r.FormValue("salary"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Env.Logger.Error("cannot parse form value 'salary'")
		return
	}

	if err = repository.CreateUser(name, age, salary, h.Env); err != nil {
		h.Env.Logger.Error("cannot create user")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) findUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Env.Logger.Error("cannot find user, wrong id format")
		return
	}

	user, err := repository.FindUserByID(id, h.Env)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Env.Logger.Error(fmt.Sprintf("cannot find user, err: %s", err.Error()))
		w.Write([]byte("No such user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.String()))
}
