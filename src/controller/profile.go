package controller

import (
	"net/http"
	"github.com/gorilla/context"
)

func Profile(tpl *TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		User string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tplVars.User = context.Get(r, "User").(string)

		tpl.Render(w, r, tplVars)
	}
}
