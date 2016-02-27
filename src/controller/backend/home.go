package backend

import (
	"net/http"
	"github.com/dfernandez/geb/src/controller"
)

func Home(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{}

	return func(w http.ResponseWriter, r *http.Request) {

		tpl.Render(w, r, tplVars)
	}
}