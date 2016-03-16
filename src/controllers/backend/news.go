package backend

import (
    "net/http"
    "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

    "github.com/dfernandez/geb/src/models/news"
    "github.com/dfernandez/geb/src/models/user"
	"github.com/dfernandez/gcore/controller"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var News = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "backend/news/news.html",
		Layout:   "backend.html",
	}

    var tplVars struct{
        News []news.News
    }

    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        tplVars.News = news.Newss(mongoSession.(*mgo.Session))

        tpl.Render(w, r, tplVars)
    }
}()

var NewsCreate = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "backend/news/create.html",
		Layout:   "backend.html",
	}

    var tplVars struct{}

    return func(w http.ResponseWriter, r *http.Request) {
        tpl.Render(w, r, tplVars)
    }
}()

var NewsEdit = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "backend/news/edit.html",
		Layout:   "backend.html",
	}

    var tplVars struct{
        News *news.News
    }

    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        vars := mux.Vars(r)

        tplVars.News = &news.News{Id: bson.ObjectIdHex(vars["id"])}
        tplVars.News.Load(mongoSession.(*mgo.Session))

        tpl.Render(w, r, tplVars)
    }
}()

var NewsSave = func() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        user := context.Get(r, "user").(user.User)
        mongoSession := context.Get(r, "mongoDB")

        r.ParseForm()
        n := news.NewNews(r.FormValue("title"), r.FormValue("htmlBody"), user)
        n.Insert(mongoSession.(*mgo.Session))

        http.Redirect(w, r, "/admin/news", http.StatusFound)
    }
}()

var NewsDelete = func() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        vars := mux.Vars(r)

        n := &news.News{Id: bson.ObjectIdHex(vars["id"])}
        n.Delete(mongoSession.(*mgo.Session))

        http.Redirect(w, r, "/admin/news", http.StatusFound)
    }
}()