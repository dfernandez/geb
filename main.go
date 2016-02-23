package main

import (
	"github.com/dfernandez/geb/server"
	"github.com/dfernandez/geb/config"
)

func main() {
	srv := server.Server{Addr: config.SrvAddr}
	srv.Boot()
}