package controller

import (
	"net/http"
	"html/template"
)

type TplVars struct {
	Title string
	Body string
}

func Home(t *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tplVars := TplVars{
			Title: "Go web!",
			Body: "Golang web application",
		}

		t.ExecuteTemplate(w, "layout", tplVars)
	}
}
