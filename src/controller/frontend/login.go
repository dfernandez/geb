package frontend

import (
	"net/http"
	"golang.org/x/oauth2"
	"github.com/gorilla/securecookie"
	"github.com/dfernandez/geb/config"
	"time"
	log "github.com/Sirupsen/logrus"
	"github.com/dfernandez/geb/src/domain"
	"github.com/gorilla/sessions"
	"github.com/dfernandez/geb/src/controller"
)

func Login(tpl *controller.TplController) func(w http.ResponseWriter, r *http.Request) {
	var tplVars struct{}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Render(w, r, tplVars)
	}
}

func OAuthLogin(conf *config.OAuthConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := conf.OAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func OAuthCallback(conf *config.OAuthConfig) func(w http.ResponseWriter, r *http.Request) {
	s := securecookie.New(config.HashKey, config.BlockKey)

	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code");
		if code == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		token, err := conf.OAuthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Error(err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if conf.Platform == "fb" {
			conf.ProfileEndpoint += token.Extra("access_token").(string)
		}

		domainToken := &domain.Token{
			OAuthToken: token,
			RefreshToken: token.RefreshToken,
			Platform: conf.Platform,
			ProfileUrl: conf.ProfileEndpoint,
		}

		encodedValue, err := s.Encode("X-Authorization", domainToken)
		if err == nil {
			// if doing a callback, delete current session profile
			store := sessions.NewCookieStore(config.HashKey)
			session, _ := store.Get(r, config.SessionName)
			session.Values["profile"] = nil
			session.Save(r, w)

			cookie := &http.Cookie{
				Name:  "X-Authorization",
				Value: encodedValue,
				Path:  "/",
				Expires: time.Now().Add(7 * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
}
