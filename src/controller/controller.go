package controller

import (
	"net/http"
	"github.com/gorilla/context"
	"html/template"
	log "github.com/Sirupsen/logrus"
)

type TplController struct {
	Template   string
	Layout     string
	Controller string
	Title      string
	Profile    map[string]interface{}
	TplVars    interface{}
}

func (tpl TplController) Render(w http.ResponseWriter, r *http.Request, tplVars interface{}) {

	t, err := template.ParseFiles(tpl.Layout, tpl.Template)
	if err != nil {
		log.Fatal(err)
	}

	tpl.Title      = "Go web!"
	tpl.Controller = r.URL.Path
	tpl.TplVars    = tplVars

	if profile := context.Get(r, "profile"); profile != nil {
		tpl.Profile = profile.(map[string]interface{})
	}

	err = t.ExecuteTemplate(w, "layout", tpl)
	if err != nil {
		log.Error(err)
	}
}