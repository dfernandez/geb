package server

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/dfernandez/geb/controller"
	"github.com/dfernandez/geb/server/decorator"
	"html/template"
)

var Router = func() *mux.Router {
	// decorators
	logger := decorator.NewLogger()

	// router
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = use(controller.Error404, logger)

	// controllers
	router.HandleFunc("/", use(controller.Home(render("view/home.html")), logger))
	router.HandleFunc("/error500", use(controller.Error500, logger))

	return router
}()

type httpHandlerDecorator interface {
	Do(http.HandlerFunc) http.HandlerFunc
}

func use(h http.HandlerFunc, decors ...httpHandlerDecorator) http.HandlerFunc {
	for _, decorator := range decors {
		h = decorator.Do(h)
	}
	return h
}

func render(tpl string) *template.Template {
	t, _ := template.ParseFiles("layout/layout.html", tpl)

	return t
}