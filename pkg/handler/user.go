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
	password := r.FormValue("password")

	if len(password) < 8 || len(username) < 8 {
		RedirectWothCookie(w, r, "Not enough length for username or password. Should be more than 8", "/signup")
		return
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		RedirectWothCookie(w, r, "Cannot create user, check data", "/signup")
		return
	}

	if err = repository.CreateUser(username, string(password_hash), h.Env); err != nil {
		RedirectWothCookie(w, r, "Cannot create user, check data", "/signup")
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handler) logIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	id, password_hash, err := repository.FindUserByUsername(username, h.Env)
	if err != nil {
		RedirectWothCookie(w, r, "Wrong username or password", "/login")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password)); err != nil {
		RedirectWothCookie(w, r, "Wrong username or password", "/login")
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
