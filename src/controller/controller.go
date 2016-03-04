package controller

import (
	"net/http"
	"github.com/gorilla/context"
	"html/template"
	log "github.com/Sirupsen/logrus"
	"github.com/dfernandez/geb/src/domain"
)

type TplController struct {
	Template   string
	Layout     string
	Controller string
	Title      string
	Profile    domain.Profile
	TplVars    interface{}
}

func (tpl TplController) Render(w http.ResponseWriter, r *http.Request, tplVars interface{}) {

	funcMap := template.FuncMap{
		"add": func(x int, y int) int {
			return x + y
		},
	}

	t := template.Must(template.New("").Funcs(funcMap).ParseFiles(tpl.Layout, tpl.Template))

	tpl.Title      = "Go web!"
	tpl.TplVars    = tplVars
	tpl.Controller = r.URL.Path

	if profile := context.Get(r, "profile"); profile != nil {
		tpl.Profile = profile.(domain.Profile)
	}

	err := t.ExecuteTemplate(w, "layout", tpl)
	if err != nil {
		log.Error(err)
	}
}