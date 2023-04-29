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

func (h *Handler) userPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(w.Header().Get("user-id"))
	h.Env.Logger.Info(fmt.Sprint(id))
	if err != nil {
		h.Env.Logger.Error("cannot parse user id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := repository.FindUserByID(id, h.Env)
	if err != nil {
		h.Env.Logger.Error("cannot find user or information related")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/userPage.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/footer.html", h.Env.Config.StaticDir),
	)
	if err != nil {
		h.Env.Logger.Error("can't parse template index.html")
		return
	}

	tmpl.ExecuteTemplate(w, "user", user)
}

func (h *Handler) logInPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/login.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/footer.html", h.Env.Config.StaticDir),
	)
	if err != nil {
		h.Env.Logger.Error("can't parse template login.html")
		return
	}

	tmpl.ExecuteTemplate(w, "login", GetCookieErrMessage(w, r, "/login"))
}

func (h *Handler) signUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/signup.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/footer.html", h.Env.Config.StaticDir),
	)
	if err != nil {
		h.Env.Logger.Error("can't parse template signup.html")
		return
	}

	tmpl.ExecuteTemplate(w, "signup", GetCookieErrMessage(w, r, "/signup"))
}

func (h *Handler) editUserInfoPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(w.Header().Get("user-id"))
	// h.Env.Logger.Info(fmt.Sprint(id))
	if err != nil {
		h.Env.Logger.Error("cannot parse user id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := repository.FindUserByID(id, h.Env)
	if err != nil {
		h.Env.Logger.Error("cannot find user or information related")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/editUserInfoPage.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/footer.html", h.Env.Config.StaticDir),
	)
	if err != nil {
		h.Env.Logger.Error("can't parse template editUserInfoPage.html")
		return
	}

	tmpl.ExecuteTemplate(w, "editUserInfo", user)
}
