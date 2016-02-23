package decorator

import (
	"net/http"
	"github.com/gorilla/securecookie"
	"log"
	"github.com/gorilla/context"
	"github.com/dfernandez/geb/config"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a Auth) Do(h http.HandlerFunc) http.HandlerFunc {
	s := securecookie.New(config.HashKey, config.BlockKey)

	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("X-Authorization")

		if (err != nil) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return;
		}

		// Redirects to auth handler
		var auth string
		err = s.Decode("X-Authorization", cookie.Value, &auth)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return;
		}

		context.Set(r, "AuthToken", auth)
		h(w, r)
	}
}
