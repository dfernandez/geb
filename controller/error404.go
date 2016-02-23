package controller

import (
	"net/http"
	"html/template"
)

func Error404(t *template.Template, tplData TplData) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "layout", tplData)
	}
}