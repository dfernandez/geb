package decorator

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/dfernandez/geb/config"
)

type Context struct {
}

func NewContext() *Context {
	return &Context{}
}

func (c Context) Do(h http.HandlerFunc) http.HandlerFunc {
	store := sessions.NewCookieStore(config.HashKey)

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, config.SessionName)
		context.Set(r, "user", session.Values["user"])

		h(w, r)
	}
}