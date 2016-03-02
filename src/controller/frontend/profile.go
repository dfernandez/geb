package frontend

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/src/controller"
)

func Profile(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct {
		Profile map[string]interface{}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tplVars.Profile = context.Get(r, "profile").(map[string]interface{})
		tpl.Render(w, r, tplVars)
	}
}
