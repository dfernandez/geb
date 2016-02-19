package server

import (
	"log"
	"net/http"
)

type Server struct {
	Addr string
}

func (s Server) Boot() {
	err := http.ListenAndServe(s.Addr, Router)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}