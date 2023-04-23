package handler

import (
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
			w.Header().Add("err", "Authorization token is empty")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if _, err := ParseToken(token); err != nil {
			w.Header().Add("err", "Token is incorrect")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// h.Env.Logger.Info(fmt.Sprintf("id: %d", userId))
		w.WriteHeader(http.StatusAccepted)

		next.ServeHTTP(w, r)
	})
}
