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
	"log"
	"encoding/gob"
)

type Session struct{}

type facebookProfilePicture struct {
	ID      string `json:"id"`
	Picture struct {
				Data struct {
						 IsSilhouette bool   `json:"is_silhouette"`
						 URL          string `json:"url"`
					 } `json:"data"`
			} `json:"picture"`
}

func NewSession() *Session {
	return &Session{}
}

func (s Session) Do(h http.HandlerFunc) http.HandlerFunc {
	store := sessions.NewCookieStore(config.HashKey)
	store.MaxAge(3600)

	ss := securecookie.New(config.HashKey, config.BlockKey)

	gob.Register(domain.Profile{})

	return func(w http.ResponseWriter, r *http.Request) {
		var profile domain.Profile
		var session, _ = store.Get(r, config.SessionName)

		if p := session.Values["profile"]; p == nil {
			cookie, err := r.Cookie("X-Authorization")

			if (err != nil) {
				h(w, r)
				return;
			}

			var token *domain.Token
			err = ss.Decode("X-Authorization", cookie.Value, &token)
			if err != nil {
				h(w, r)
				return;
			}

			var conf *oauth2.Config
			if token.Platform == "g+" {
				conf = config.GoogleOAuthConfig
			} else {
				conf = config.FacebookOAuthConfig
			}

			client := conf.Client(oauth2.NoContext, token.OAuthToken)
			response, err := client.Get(token.ProfileUrl)

			if err != nil {
				log.Println(err)
				h(w, r)
				return;
			}

			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)

			json.Unmarshal(body, &profile)

			// Facebook profile image
			if token.Platform == "fb" {
				response, _ = client.Get("https://graph.facebook.com/" + profile.ID + "?fields=picture.type(large)")
				defer response.Body.Close()
				body, _ = ioutil.ReadAll(response.Body)

				var profilePicture facebookProfilePicture
				json.Unmarshal(body, &profilePicture)

				profile.Picture = profilePicture.Picture.Data.URL
			}

			session.Values["profile"] = profile
			session.Save(r, w)

			err = store.Save(r, w, session)
			if err != nil {
				log.Println(err)
			}

			profile.UpdateActivity()
		} else {
			profile = p.(domain.Profile)
		}

		context.Set(r, "profile", profile)

		h(w, r)
	}
}