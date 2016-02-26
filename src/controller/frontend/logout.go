package frontend

import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/dfernandez/geb/config"
)

func Logout() func(w http.ResponseWriter, r *http.Request) {
	store := sessions.NewCookieStore(config.HashKey)

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, config.SessionName)

		session.Values["profile"] = nil
		session.Save(r, w)

		cookie := &http.Cookie{
			Name:   "X-Authorization",
			Path:   "/",
			MaxAge: -1,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
