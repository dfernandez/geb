package server

import (
	"log"
	"net/http"
	"github.com/dfernandez/geb/controllers"
	"github.com/dfernandez/geb/server/decorator"
	"os"
	"github.com/gorilla/mux"
)

type Server struct {
	Addr string
}

func New() *Server {
	s := Server{}
	s.init()

	return &s
}

func (s Server) init() {
	// decorators
	logger200 := decorator.Logger{log.New(os.Stdout, "200: ", log.LstdFlags)}
	logger404 := decorator.Logger{log.New(os.Stdout, "404: ", log.LstdFlags)}

	// routes
	r := mux.NewRouter()
	r.HandleFunc("/", decorator.Decorate(controllers.Home, logger200))

	// 404 not found
	r.NotFoundHandler = http.HandlerFunc(decorator.Decorate(controllers.NotFound, logger404))

	http.Handle("/", r)
}

func (s Server) Boot() {
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}