package controller

import (
	"net/http"
)

func Logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:   "X-Authorization",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
