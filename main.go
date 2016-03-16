package main

import (
	"github.com/dfernandez/gcore"
	"github.com/dfernandez/gcore/decorator"
	"github.com/dfernandez/geb/config"
	"github.com/dfernandez/geb/src/controllers/backend"
	"github.com/dfernandez/geb/src/controllers/frontend"
	"github.com/dfernandez/geb/src/models/user"
)

func main() {
    gcore := gcore.New(config.SrvAddr)

	gcore.RegisterDecorator(decorator.NewContext(config.HashKey, config.SessionName))
	gcore.RegisterDecorator(decorator.NewLogger())
	gcore.RegisterDecorator(decorator.NewRecover())

	auth    := decorator.NewAuth()
	admin   := decorator.NewAdmin(user.IsAdmin)
	mongodb := decorator.NewMongo(config.MongoServer)

	gcore.AddStatic("./public", "/")
	gcore.AddStatic("./public", "/js/")
	gcore.AddStatic("./public", "/css/")
	gcore.AddStatic("./public", "/fonts/")

	gcore.AddRoute("/",               frontend.Home)
	gcore.AddRoute("/login",          frontend.Login)
	gcore.AddRoute("/login/callback", frontend.Callback, mongodb)
	gcore.AddRoute("/logout",         frontend.Logout)
	gcore.AddRoute("/profile",        frontend.Profile, auth)

	gcore.AddRoute("/admin",                  backend.Home, admin, mongodb)
	gcore.AddRoute("/admin/news",             backend.News, admin, mongodb)
	gcore.AddRoute("/admin/news/create",      backend.NewsCreate, admin, mongodb)
	gcore.AddRoute("/admin/news/save",        backend.NewsSave, admin, mongodb)
	gcore.AddRoute("/admin/news/edit/{id}",   backend.NewsEdit, admin, mongodb)
	gcore.AddRoute("/admin/news/delete/{id}", backend.NewsDelete, admin, mongodb)
	gcore.AddRoute("/admin/users",            backend.Users, admin, mongodb)

	gcore.Boot()
}