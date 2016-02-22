package server

import (
	"github.com/gorilla/mux"

	"github.com/dfernandez/geb/controller"
	"github.com/dfernandez/geb/server/decorator"
)

var Router = func() *mux.Router {
	// decorators
	logger := decorator.NewLogger()

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = useHandler(controller.Error404, logger)

	// controllers
	router.HandleFunc("/", useHandler(controller.Home(useTemplate("view/home.html")), logger))

	// errors
	router.HandleFunc("/error500", useHandler(controller.Error500, logger))

	return router
}()