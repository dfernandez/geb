package server

import (
	"github.com/gorilla/mux"

	"github.com/dfernandez/geb/controller"
	"github.com/dfernandez/geb/server/decorator"
	"net/http"
	"strings"
)

var Router = func() *mux.Router {
	// decorators
	logger := decorator.NewLogger()
	auth   := decorator.NewAuth()

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = useHandler(controller.Error404(useTemplate("templates/error404.html")), logger)

	// static files
	router.HandleFunc("/{file}", useHandler(serveStatics, logger)).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return strings.Contains(r.RequestURI, ".")
	})

	// controllers
	router.HandleFunc("/",        useHandler(controller.Home(useTemplate("templates/home.html")),       logger))
	router.HandleFunc("/login",   useHandler(controller.Login(useTemplate("templates/login.html")),     logger))
	router.HandleFunc("/logout",  useHandler(controller.Logout(),                                       logger))
	router.HandleFunc("/private", useHandler(controller.Private(useTemplate("templates/private.html")), auth, logger))

	return router
}()

func serveStatics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	http.ServeFile(w, r, "./public/" + file)
}