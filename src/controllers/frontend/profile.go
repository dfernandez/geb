package frontend

import (
    "net/http"
    "github.com/gorilla/context"
    "github.com/dfernandez/geb/src/models/user"
	"github.com/dfernandez/gcore/controller"
)

var Profile = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "frontend/profile/profile.html",
		Layout:   "frontend.html",
	}

    var tplVars struct {
        Profile *user.User
    }

    return func(w http.ResponseWriter, r *http.Request) {
        u := context.Get(r, "user").(user.User)
        tplVars.Profile = &u
        tpl.Render(w, r, tplVars)
    }
}()
