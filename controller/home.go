package controller

import (
	"net/http"
	"html/template"
)

func Home(t *template.Template, tplVars TplVars) func(w http.ResponseWriter, r *http.Request) {
	var homeVars struct {
		Text1 string
		Text2 string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		homeVars.Text1 = "text 1"
		homeVars.Text2 = "text 2"

		tplVars.Body = homeVars
		t.ExecuteTemplate(w, "layout", tplVars)
	}
}
