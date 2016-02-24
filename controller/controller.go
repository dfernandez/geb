package controller

import (
	"net/http"
	"github.com/gorilla/context"
	"html/template"
	"log"
)

type TplController struct {
	Template   string
	Controller string
	Title      string
	User       string
	TplVars    interface{}
}

func (tpl TplController) Render(w http.ResponseWriter, r *http.Request, tplVars interface{}) {

	t, err := template.ParseFiles("layout/layout.html", tpl.Template)
	if err != nil {
		log.Fatal(err)
	}

	tpl.Title      = "Go web!"
	tpl.Controller = r.URL.Path
	tpl.TplVars    = tplVars

	if user := context.Get(r, "User"); user != nil {
		tpl.User = user.(string)
	}

	t.ExecuteTemplate(w, "layout", tpl)
}