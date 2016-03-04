package backend

import (
    "net/http"
    "github.com/dfernandez/geb/src/controller"
    "gopkg.in/mgo.v2"
    "github.com/gorilla/context"
    "github.com/dfernandez/geb/src/models/news"
    "github.com/dfernandez/geb/src/models/user"
    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2/bson"
)

func News(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
    var tplVars struct{
        News []news.News
    }

    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        tplVars.News = news.Newss(mongoSession.(*mgo.Session))

        tpl.Render(w, r, tplVars)
    }
}

func NewsCreate(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
    var tplVars struct{}

    return func(w http.ResponseWriter, r *http.Request) {
        tpl.Render(w, r, tplVars)
    }
}

func NewsSave() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        user := context.Get(r, "user").(user.User)
        mongoSession := context.Get(r, "mongoDB")

        r.ParseForm()
        n := news.NewNews(r.FormValue("title"), r.FormValue("body"), user)
        n.Insert(mongoSession.(*mgo.Session))

        http.Redirect(w, r, "/admin/news", http.StatusFound)
    }
}

func NewsDelete() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        mongoSession := context.Get(r, "mongoDB")
        vars := mux.Vars(r)

        n := &news.News{Id: bson.ObjectIdHex(vars["id"])}
        n.Delete(mongoSession.(*mgo.Session))

        http.Redirect(w, r, "/admin/news", http.StatusFound)
    }
}