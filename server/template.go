package server

import (
	"html/template"
	"github.com/dfernandez/geb/controller"
)

func useTemplate(tpl string) (*template.Template, controller.TplVars) {

	t, _ := template.ParseFiles("layout/layout.html", tpl)

	tplVars := controller.TplVars{
		Title: "Go web!",
	}

	return t, tplVars
}
