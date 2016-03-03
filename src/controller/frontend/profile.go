package frontend

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/src/controller"
	"github.com/dfernandez/geb/src/domain"
)

func Profile(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		Profile *domain.Profile
	}

	return func(w http.ResponseWriter, r *http.Request) {
		p := context.Get(r, "profile").(domain.Profile)
		tplVars.Profile = &p
		tpl.Render(w, r, tplVars)
	}
}
