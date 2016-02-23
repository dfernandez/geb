package server

import (
	"log"
	"html/template"
	"github.com/dfernandez/geb/controller"
)

func useTemplate(tpl string) (*template.Template, controller.TplData) {

	t, err := template.ParseFiles("layout/layout.html", tpl)
	if err != nil {
		log.Fatal(err)
	}

	tplData := controller.TplData{
		Title: "Go web!",
	}

	return t, tplData
}
