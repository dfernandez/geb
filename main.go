package main

import (
	"github.com/dfernandez/geb/server"
)

func main() {
	srv := server.New()
	srv.Boot()
}