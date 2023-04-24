package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt-token")
		if err != nil {
			w.Header().Add("err", "Cannot get cookie jwt-token")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		token := cookie.Value
		if token == "" {
			RedirectWothCookie(w, r, "You are not authorized", "/login")
			return
		}

		userId, err := ParseToken(token)
		if err != nil {
			RedirectWothCookie(w, r, "Wrong username or password", "/login")
			return
		}

		// h.Env.Logger.Info(fmt.Sprintf("id: %d", userId))
		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("user-id", fmt.Sprint(userId))
		next.ServeHTTP(w, r)
	})
}
