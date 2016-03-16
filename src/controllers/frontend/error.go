package frontend

import (
    "net/http"
	"github.com/dfernandez/gcore/controller"
)

var Error404 = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "frontend/error/error404.html",
		Layout:   "frontend.html",
	}
    var tplVars struct{}

    return func(w http.ResponseWriter, r *http.Request) {
        tpl.Render(w, r, tplVars)
    }
}()