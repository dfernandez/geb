package controller

import (
	"net/http"
	"github.com/dfernandez/geb/src/domain"
)

func Admin(tpl *TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{}

	return func(w http.ResponseWriter, r *http.Request) {

		tpl.Render(w, r, tplVars)
	}
}

func AdminUsers(tpl *TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{
		Profiles []domain.Profile
	}

	return func(w http.ResponseWriter, r *http.Request) {
		p := &domain.Profile{}
		tplVars.Profiles = p.GetProfiles()

		tpl.Render(w, r, tplVars)
	}
}
