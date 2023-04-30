package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/NikiTesla/geant4help/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(password) < 6 || len(username) < 6 {
		RedirectWithCookie(w, r, "Not enough length for username or password. Should be more than 6", "/signup")
		return
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		RedirectWithCookie(w, r, "Cannot create user, check data", "/signup")
		return
	}

	if err = repository.CreateUser(username, string(password_hash), h.Env); err != nil {
		RedirectWithCookie(w, r, "User already exists", "/signup")
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handler) logIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	id, password_hash, err := repository.FindUserByUsername(username, h.Env)
	if err != nil {
		RedirectWithCookie(w, r, "Wrong username or password", "/login")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password)); err != nil {
		RedirectWithCookie(w, r, "Wrong username or password", "/login")
		return
	}

	token, err := GenerateToken(id)
	if err != nil {
		h.Env.Logger.Error("cannot generate token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// h.Env.Logger.Info(token)
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt-token",
		Value: token,
		Path:  "/",
	})

	http.Redirect(w, r, "/user/", http.StatusSeeOther)
}

func (h *Handler) editUserInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(w.Header().Get("user-id"))
	if err != nil {
		RedirectWithCookie(w, r, err.Error(), "/user")
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		RedirectWithCookie(w, r, err.Error(), "/user")
		return
	}

	job := r.FormValue("job")

	err = repository.EditUserInfo(id, name, email, age, job, h.Env)
	if err != nil {
		RedirectWithCookie(w, r, err.Error(), "/user")
		return
	}

	h.userPage(w, r)
}

func (h *Handler) signOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt-token",
		Value:    "fbsdfsd",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	h.indexPage(w, r)
}
