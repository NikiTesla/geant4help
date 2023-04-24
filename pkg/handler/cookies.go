package handler

import "net/http"

// GetCookieErrMessage parse authorization error from cookie if exists
// then empties it for path
// TODO how to rewrite it correctly ?
func GetCookieErrMessage(w http.ResponseWriter, r *http.Request, path string) string {
	var msg string
	cookie, err := r.Cookie("Err")
	if err == nil {
		msg = cookie.Value
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Err",
		Value: "",
		Path:  path,
	})

	return msg
}

// RedirectWothCookie redirects request to url with cookie Err: msg
func RedirectWothCookie(w http.ResponseWriter, r *http.Request, msg, url string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Err",
		Value: msg,
		Path:  url,
	})

	http.Redirect(w, r, url, http.StatusSeeOther)
}
