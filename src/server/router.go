package server

import (
	"github.com/gorilla/mux"
	"github.com/dfernandez/geb/src/controller"
	"github.com/dfernandez/geb/src/server/decorator"
	"net/http"
	"strings"
	"github.com/dfernandez/geb/config"
)

var Router = func() *mux.Router {
	// decorators
	auth    := decorator.NewAuth()
	logger  := decorator.NewLogger()
	session := decorator.NewSession()

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = useHandler(controller.Error404(useTemplate("error404.html")), logger)

	// static files
	router.HandleFunc("/{file}", useHandler(serveStatics, logger)).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return strings.Contains(r.RequestURI, ".")
	})

	// home controller
	router.HandleFunc("/", useHandler(controller.Home(useTemplate("home.html")), session, logger))

	// login controller
	router.HandleFunc("/login",                   useHandler(controller.Login(useTemplate("login.html")),     session, logger))
	router.HandleFunc("/login/google",            useHandler(controller.OAuthLogin(config.GoogleConfig),      session, logger))
	router.HandleFunc("/login/google/callback",   useHandler(controller.OAuthCallback(config.GoogleConfig),   session, logger))
	router.HandleFunc("/login/facebook",          useHandler(controller.OAuthLogin(config.FacebookConfig),    session, logger))
	router.HandleFunc("/login/facebook/callback", useHandler(controller.OAuthCallback(config.FacebookConfig), session, logger))

	// logout controller
	router.HandleFunc("/logout", useHandler(controller.Logout(), session, logger))

	// profile controller
	router.HandleFunc("/profile", useHandler(controller.Profile(useTemplate("profile.html")), auth, session, logger))

	return router
}()

func serveStatics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	http.ServeFile(w, r, "./public/" + file)
}