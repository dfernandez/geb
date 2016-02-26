package backend

import (
	"net/http"
	"github.com/dfernandez/geb/src/domain"
	"github.com/dfernandez/geb/src/controller"
)

func Users(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{
		Profiles []domain.Profile
	}

	return func(w http.ResponseWriter, r *http.Request) {
		p := &domain.Profile{}
		tplVars.Profiles = p.GetProfiles()

		tpl.Render(w, r, tplVars)
	}
}