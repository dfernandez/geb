package backend

import (
	"net/http"
	"github.com/dfernandez/geb/src/domain"
	"github.com/dfernandez/geb/src/controller"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/context"
)

func Users(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{
		Profiles []domain.Profile
	}

	return func(w http.ResponseWriter, r *http.Request) {
		mongoSession := context.Get(r, "mongoDB")
		p := &domain.Profile{}
		tplVars.Profiles = p.GetProfiles(mongoSession.(*mgo.Session))

		tpl.Render(w, r, tplVars)
	}
}