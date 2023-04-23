package handler

import (
	"net/http"

	"github.com/NikiTesla/geant4help/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password_hash, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.Env.Logger.Error("cannot generate hash from password")
		return
	}

	if err = repository.CreateUser(username, string(password_hash), h.Env); err != nil {
		h.Env.Logger.Error("cannot create user")
		return
	}

	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
}

func (h *Handler) logIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	id, password_hash, err := repository.FindUserByUsername(username, h.Env)
	if err != nil {
		w.Header().Add("err", "Can't find user with this username")
		r.Method = "GET"
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password)); err != nil {
		w.Header().Add("err", "Wrong password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	token, err := GenerateToken(id)
	if err != nil {
		h.Env.Logger.Error("cannot generate token")
		w.WriteHeader(http.StatusInternalServerError)
	}

	// h.Env.Logger.Info(token)
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt-token",
		Value: token,
		Path:  "/",
	})

	http.Redirect(w, r, "/user/", http.StatusSeeOther)
}
