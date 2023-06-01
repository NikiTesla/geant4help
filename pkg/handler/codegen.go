package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	codegen "github.com/NikiTesla/geant4help/pkg/code_gen"
)

func (h *Handler) codeGenPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("%v/html/codeGenPage.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/user_header.html", h.Env.Config.StaticDir),
		fmt.Sprintf("%v/html/user_footer.html", h.Env.Config.StaticDir),
	)

	if err != nil {
		h.Env.Logger.Error("can't parse template codeGenPage")
		return
	}

	tmpl.ExecuteTemplate(w, "codeGenPage", nil)
}

func (h *Handler) codeGen(w http.ResponseWriter, r *http.Request) {
	shape := r.FormValue("shapes")
	// TODO rewrite it
	sizeX, _ := strconv.ParseFloat(r.FormValue("size_x"), 64)
	sizeY, _ := strconv.ParseFloat(r.FormValue("size_y"), 64)
	sizeZ, _ := strconv.ParseFloat(r.FormValue("size_z"), 64)
	innerR, _ := strconv.ParseFloat(r.FormValue("inner_r"), 64)
	outerR, _ := strconv.ParseFloat(r.FormValue("outer_r"), 64)
	material := r.FormValue("material")

	geometry := &codegen.Geometry{
		Shape:       shape,
		SizeX:       sizeX,
		SizeY:       sizeY,
		SizeZ:       sizeZ,
		InnerRadius: innerR,
		OuterRadius: outerR,
		Material:    material,
	}

	err := codegen.GenerateMacroFile(geometry)
	if err != nil {
		h.Env.Logger.Error("cannot generate macro file")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot generate macro file"))
		return
	}

	w.Write([]byte("Successfully generated!"))
}
