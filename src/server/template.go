package server

import (
	"github.com/dfernandez/geb/src/controller"
)

func useTemplate(tpl string) (* controller.TplController) {
	return &controller.TplController{
		Template: "src/templates/" + tpl,
		Layout:   "src/layout/layout.html",
	}
}
