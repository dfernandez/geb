package backend

import (
    "net/http"
    "github.com/dfernandez/geb/src/controller"
    "gopkg.in/mgo.v2"
    "github.com/gorilla/context"
    "github.com/dfernandez/geb/src/models/user"
)

func Users(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
    var tplVars struct{
        Profiles []user.User
    }

    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        tplVars.Profiles = user.Users(mongoSession.(*mgo.Session))

        tpl.Render(w, r, tplVars)
    }
}