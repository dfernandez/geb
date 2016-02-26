package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"github.com/dfernandez/geb/config"
	"github.com/dfernandez/geb/src/server/decorator"
	"github.com/dfernandez/geb/src/controller/backend"
	"github.com/dfernandez/geb/src/controller/frontend"
)

var Router = func() *mux.Router {
	// decorators
	admin   := decorator.NewAdmin()
	auth    := decorator.NewAuth()
	logger  := decorator.NewLogger()
	session := decorator.NewSession()

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = useHandler(frontend.Error404(useTemplate("error404.html")), logger)

	// static files
	router.HandleFunc("/{file}", useHandler(serveStatics, logger)).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return strings.Contains(r.RequestURI, ".")
	})

	// home controller
	router.HandleFunc("/", useHandler(frontend.Home(useTemplate("home.html")), session, logger))

	// login controller
	router.HandleFunc("/login",                   useHandler(frontend.Login(useTemplate("login.html")),     session, logger))
	router.HandleFunc("/login/google",            useHandler(frontend.OAuthLogin(config.GoogleConfig),      session, logger))
	router.HandleFunc("/login/google/callback",   useHandler(frontend.OAuthCallback(config.GoogleConfig),   session, logger))
	router.HandleFunc("/login/facebook",          useHandler(frontend.OAuthLogin(config.FacebookConfig),    session, logger))
	router.HandleFunc("/login/facebook/callback", useHandler(frontend.OAuthCallback(config.FacebookConfig), session, logger))

	// logout controller
	router.HandleFunc("/logout", useHandler(frontend.Logout(), session, logger))

	// profile controller
	router.HandleFunc("/profile", useHandler(frontend.Profile(useTemplate("profile.html")), auth, session, logger))

	// admin
	router.HandleFunc("/admin",       useHandler(backend.Home(useBackendTemplate("home.html")),   admin, session, logger))
	router.HandleFunc("/admin/users", useHandler(backend.Users(useBackendTemplate("users.html")), admin, session, logger))

	return router
}()

func serveStatics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	http.ServeFile(w, r, "./public/" + file)
}