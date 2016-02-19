package main

import (
	"github.com/dfernandez/geb/server"
)

func main() {
	srv := server.Server{Addr:":8000"}
	srv.Boot()
}