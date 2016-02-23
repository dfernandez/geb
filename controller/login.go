package controller

import (
	"net/http"
	"html/template"
	"github.com/gorilla/securecookie"
	"github.com/dfernandez/geb/config"
)

func Login(t *template.Template, tplData TplData) func(w http.ResponseWriter, r *http.Request) {
	s := securecookie.New(config.HashKey, config.BlockKey)

	return func(w http.ResponseWriter, r *http.Request) {
		encodedValue, err := s.Encode("X-Authorization", "qwerty123456")
		if err == nil {
			cookie := &http.Cookie{
				Name:  "X-Authorization",
				Value: encodedValue,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}

		t.ExecuteTemplate(w, "layout", tplData)
	}
}
