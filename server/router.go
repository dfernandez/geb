package server

import (
	"github.com/gorilla/mux"

	"github.com/dfernandez/geb/controller"
	"github.com/dfernandez/geb/server/decorator"
)

var Router = func() *mux.Router {
	// decorators
	logger := decorator.NewLogger()
	auth   := decorator.NewAuth()

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = useHandler(controller.Error404(useTemplate("templates/error404.html")), logger)

	// controllers
	router.HandleFunc("/", useHandler(controller.Home(useTemplate("templates/home.html")), auth, logger))

	return router
}()