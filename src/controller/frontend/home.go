package frontend

import (
    "net/http"
    "github.com/dfernandez/geb/src/controller"
)

func Home(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
    var tplVars struct {
        Text1 string
        Text2 string
    }

    return func(w http.ResponseWriter, r *http.Request) {
        tplVars.Text1 = "text 1"
        tplVars.Text2 = "text 2"

        tpl.Render(w, r, tplVars)
    }
}
