package server

import (
    "github.com/dfernandez/geb/src/controller"
)

func useTemplate(tpl string) (* controller.TplController) {
    return &controller.TplController{
        Template: "src/templates/frontend/" + tpl,
        Layout:   "src/templates/frontend.html",
    }
}

func useBackendTemplate(tpl string) (* controller.TplController) {
    return &controller.TplController{
        Template: "src/templates/backend/" + tpl,
        Layout:   "src/templates/backend.html",
    }
}
