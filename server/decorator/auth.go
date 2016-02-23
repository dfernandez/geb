package decorator

import (
	"net/http"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a Auth) Do(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
