package controller

import (
	"net/http"
	"html/template"
	"github.com/gorilla/context"
)

func Private(t *template.Template, tplData TplData) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		AuthToken string
	}

	tplData.Controller = "private"

	return func(w http.ResponseWriter, r *http.Request) {
		tplVars.AuthToken = context.Get(r, "AuthToken").(string)

		tplData.TplVars = tplVars
		t.ExecuteTemplate(w, "layout", tplData)
	}
}
