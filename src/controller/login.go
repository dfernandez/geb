package controller

import (
	"net/http"
	"golang.org/x/oauth2"
	"github.com/gorilla/securecookie"
	"github.com/dfernandez/geb/config"
	"io/ioutil"
	"encoding/json"
	"github.com/dfernandez/geb/src/domain"
)

func Login(tpl *TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w, r, tplVars)
	}
}

func OAuthLogin(conf *oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func GoogleCallback(conf *oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	s := securecookie.New(config.HashKey, config.BlockKey)

	return func(w http.ResponseWriter, r *http.Request) {
		token, _    := conf.Exchange(oauth2.NoContext, r.URL.Query().Get("code"))
		client      := conf.Client(oauth2.NoContext, token)
		response, _ := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		var profile domain.Profile
		json.Unmarshal(body, &profile)

		encodedValue, err := s.Encode("X-Authorization", profile.Email)
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
