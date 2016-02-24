package controller

import (
	"net/http"
	"github.com/dfernandez/geb/src/domain"
	"github.com/gorilla/context"
)

func Profile(tpl *TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		Profile domain.Profile
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tplVars.Profile = context.Get(r, "profile").(domain.Profile)
		tpl.Render(w, r, tplVars)
	}
}
