package server

import (
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Addr string
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (s Server) Boot() {
	srv := &http.Server{
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on... ", s.Addr)

	err = srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
	if err != nil {
		log.Fatal(err)
	}
}