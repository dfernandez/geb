package decorator

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/src/models/user"
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
		u := context.Get(r, "user")
		if u == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user := u.(user.User)
		if user.IsAdmin() {
			h(w,r)
			return
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
