package decorator

import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/dfernandez/geb/config"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/context"
)

type Session struct{}

func NewSession() *Session {
	return &Session{}
}

func (s Session) Do(h http.HandlerFunc) http.HandlerFunc {
	store := sessions.NewCookieStore(config.HashKey)
	ss    := securecookie.New(config.HashKey, config.BlockKey)

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, config.SessionName)

		if authorization := session.Values[config.SessionUser]; authorization == nil {
			cookie, err := r.Cookie("X-Authorization")

			if (err != nil) {
				h(w, r)
				return;
			}

			var auth string
			err = ss.Decode("X-Authorization", cookie.Value, &auth)
			if err != nil {
				h(w, r)
				return;
			}
			session.Values[config.SessionUser] = auth
			session.Save(r, w)
			context.Set(r, "User", auth)
		} else {
			context.Set(r, "User", authorization)
		}

		h(w, r)
	}
}