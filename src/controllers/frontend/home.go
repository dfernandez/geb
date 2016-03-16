package frontend

import (
    "net/http"
	"github.com/dfernandez/gcore/controller"
)

var Home = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "frontend/home/home.html",
		Layout:   "frontend.html",
	}

    var tplVars struct {}

    return func(w http.ResponseWriter, r *http.Request) {
        tpl.Render(w, r, tplVars)
    }
}()