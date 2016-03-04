package frontend

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/src/controller"
	"github.com/dfernandez/geb/src/models/user"
)

func Profile(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		Profile *user.User
	}

	return func(w http.ResponseWriter, r *http.Request) {
		u := context.Get(r, "user").(user.User)
		tplVars.Profile = &u
		tpl.Render(w, r, tplVars)
	}
}
