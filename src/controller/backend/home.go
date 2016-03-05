package backend

import (
    "net/http"
    "github.com/dfernandez/geb/src/controller"
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2"
    "github.com/dfernandez/geb/src/models/user"
    "github.com/dfernandez/geb/src/models/news"
)

func Home(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
    var tplVars struct{
        ProfilesCount int
        NewsCount int
    }

    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        tplVars.ProfilesCount = user.Count(mongoSession.(*mgo.Session))
        tplVars.NewsCount     = news.Count(mongoSession.(*mgo.Session))

        tpl.Render(w, r, tplVars)
    }
}
