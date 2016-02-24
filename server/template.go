package server

import (
	"github.com/dfernandez/geb/controller"
)

func useTemplate(tpl string) (* controller.TplController) {
	return &controller.TplController{Template: tpl}
}
