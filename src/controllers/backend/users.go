package backend

import (
    "net/http"
    "gopkg.in/mgo.v2"

    "github.com/gorilla/context"
    "github.com/dfernandez/geb/src/models/user"
	"github.com/dfernandez/gcore/controller"
)

var Users = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "backend/users/users.html",
		Layout:   "backend.html",
	}

    var tplVars struct{
        Profiles []user.User
    }

    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        tplVars.Profiles = user.Users(mongoSession.(*mgo.Session))

        tpl.Render(w, r, tplVars)
    }
}()