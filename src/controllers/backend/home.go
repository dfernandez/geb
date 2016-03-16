package backend

import (
    "net/http"
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2"
    "github.com/dfernandez/geb/src/models/user"
    "github.com/dfernandez/geb/src/models/news"
	"github.com/dfernandez/gcore/controller"
)

var Home = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "backend/home/home.html",
		Layout:   "backend.html",
	}

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
}()
