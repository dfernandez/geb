package backend

import (
	"net/http"
	"github.com/dfernandez/geb/src/controller"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/src/domain"
	"gopkg.in/mgo.v2"
)

func Home(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{
		ProfilesCount int
	}

	return func(w http.ResponseWriter, r *http.Request) {
		mongoSession := context.Get(r, "mongoDB")
		tplVars.ProfilesCount = domain.Count(mongoSession.(*mgo.Session))

		tpl.Render(w, r, tplVars)
	}
}
