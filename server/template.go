package server

import (
	"html/template"
	"github.com/dfernandez/geb/controller"
)

func useTemplate(tpl string) (*template.Template, controller.TplData) {

	t, _ := template.ParseFiles("layout/layout.html", tpl)

	tplData := controller.TplData{
		Title: "Go web!",
	}

	return t, tplData
}
