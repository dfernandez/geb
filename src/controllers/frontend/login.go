package frontend

import (
	log "github.com/Sirupsen/logrus"

    "net/http"
    "golang.org/x/oauth2"
    "io/ioutil"
    "encoding/json"
    "encoding/gob"

	"github.com/dfernandez/geb/config"
	"github.com/gorilla/sessions"
    "github.com/dfernandez/geb/src/models/user"
	"github.com/dfernandez/gcore/controller"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
)

var Login = func() func(w http.ResponseWriter, r *http.Request) {
	tpl := &controller.Controller{
		Template: "frontend/login/login.html",
		Layout:   "frontend.html",
	}

    var tplVars struct{
        ClientId string
        Domain string
        CallbackURL string
    }

    tplVars.ClientId    = config.OAuthConfig.ClientID
    tplVars.Domain      = "web83-es.eu.auth0.com"
    tplVars.CallbackURL = config.OAuthConfig.RedirectURL

    return func(w http.ResponseWriter, r *http.Request) {
        tpl.Render(w, r, tplVars)
    }
}()

var Logout = func() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie := &http.Cookie{
            Name:   config.SessionName,
            Path:   "/",
            MaxAge: -1,
        }

        http.SetCookie(w, cookie)
        http.Redirect(w, r, "/", http.StatusFound)
    }
}()

var Callback = func() func(w http.ResponseWriter, r *http.Request) {
    conf  := config.OAuthConfig
    store := sessions.NewCookieStore(config.HashKey)
    store.MaxAge(0)

    gob.Register(user.User{})

    return func(w http.ResponseWriter, r *http.Request) {
        // Getting the Code that we got from Auth0
        code := r.URL.Query().Get("code")

        // Exchanging the code for a token
        token, err := conf.Exchange(oauth2.NoContext, code)
        if err != nil {
            log.Error(err)
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }

        // Getting now the User information
        client := conf.Client(oauth2.NoContext, token)
        resp, err := client.Get("https://web83-es.eu.auth0.com/userinfo")
        if err != nil {
            log.Error(err)
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }

        // Reading the body
        raw, err := ioutil.ReadAll(resp.Body)
        defer resp.Body.Close()
        if err != nil {
            log.Error(err)
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }

        // Unmarshalling the JSON of the Profile
        var profile map[string]interface{}
        if err := json.Unmarshal(raw, &profile); err != nil {
            log.Error(err)
            http.Redirect(w, r, "/login", http.StatusFound)
            return
        }

        // Facebook profile picture fix
        identities := profile["identities"].([]interface{})[0].(map[string]interface{})
        if identities["provider"] == "facebook" {
            profile["picture"] = "https://graph.facebook.com/" + identities["user_id"].(string) + "/picture?width=100&height=100"
        }

        mongoSession := context.Get(r, "mongoDB")

        // User profile
        u := user.NewUser(
            profile["name"].(string),
            profile["email"].(string),
            profile["locale"].(string),
            profile["picture"].(string))

        u.Init(mongoSession.(*mgo.Session))

        // Saving the information to the session.
        var session, _ = store.Get(r, config.SessionName)
        session.Options.MaxAge         = 0
        session.Values["user"]      = u

        session.Save(r, w)
        err = store.Save(r, w, session)
        if err != nil {
            log.Error(err)
        }

        // Update last login
        u.UpdateActivity(mongoSession.(*mgo.Session))


        // Redirect to logged in page
        http.Redirect(w, r, "/profile", http.StatusMovedPermanently)
    }
}()
