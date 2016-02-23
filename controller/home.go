package controller

import (
	"net/http"
	"html/template"
)

func Home(t *template.Template, tplData TplData) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		Text1 string
		Text2 string
	}

	tplData.Controller = "home"

	return func(w http.ResponseWriter, r *http.Request) {
		tplVars.Text1 = "text 1"
		tplVars.Text2 = "text 2"

		tplData.TplVars = tplVars
		t.ExecuteTemplate(w, "layout", tplData)
	}
}
