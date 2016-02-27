package decorator

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/src/domain"
	"github.com/dfernandez/geb/config"
)

type Admin struct {
	Administrators []string
}

func NewAdmin() *Admin {
	return &Admin{Administrators: config.Administrators}
}

func (a Admin) Do(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		profile := context.Get(r, "profile")
		if profile == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		p := profile.(domain.Profile)
		if p.IsAdmin() {
			h(w,r)
			return
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
