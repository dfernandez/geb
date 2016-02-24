package controller

import (
	"net/http"
	"golang.org/x/oauth2"
	"github.com/gorilla/securecookie"
	"github.com/dfernandez/geb/config"
)

func Login(tpl *TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w, r, tplVars)
	}
}

func OAuthLogin(conf *oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := conf.AuthCodeURL("state")
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func GoogleCallback(conf *oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	s := securecookie.New(config.HashKey, config.BlockKey)

	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code");
		if code == "" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		token, _ := conf.Exchange(oauth2.NoContext, code)
		encodedValue, err := s.Encode("X-Authorization", token)
		if err == nil {
			cookie := &http.Cookie{
				Name:  "X-Authorization",
				Value: encodedValue,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
}
