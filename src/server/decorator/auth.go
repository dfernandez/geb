package decorator

import (
	"net/http"
	"github.com/gorilla/context"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a Auth) Do(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if p := context.Get(r, "profile"); p == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		h(w, r)
	}
}
