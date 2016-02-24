package decorator

import (
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/dfernandez/geb/config"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/context"
	"golang.org/x/oauth2"
	"io/ioutil"
	"github.com/dfernandez/geb/src/domain"
	"encoding/json"
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

		var profile domain.Profile
		if p := session.Values["profile"]; p == nil {
			cookie, err := r.Cookie("X-Authorization")

			if (err != nil) {
				h(w, r)
				return;
			}

			var token *oauth2.Token
			err = ss.Decode("X-Authorization", cookie.Value, &token)
			if err != nil {
				h(w, r)
				return;
			}

			conf := config.GoogleOAuthConfig

			client      := conf.Client(oauth2.NoContext, token)
			response, _ := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)

			json.Unmarshal(body, &profile)

			session.Values["profile"] = profile
			session.Save(r, w)
		}

		context.Set(r, "profile", profile)

		h(w, r)
	}
}