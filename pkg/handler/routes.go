package handler

import (
	"fmt"
	"html/template"
	"net/http"
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
