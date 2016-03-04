package server

import (
    "github.com/gorilla/mux"
    "net/http"
    "strings"
    "github.com/dfernandez/geb/src/server/decorator"
    "github.com/dfernandez/geb/src/controller/backend"
    "github.com/dfernandez/geb/src/controller/frontend"
)

var Router = func() *mux.Router {
    // decorators
    admin := decorator.NewAdmin()
    auth  := decorator.NewAuth()
    mongo := decorator.NewMongo()

    // router
    router := mux.NewRouter().StrictSlash(true)
    router.NotFoundHandler = useHandler(frontend.Error404(useTemplate("error404.html")))

    // static files
    router.HandleFunc("/{file}", useHandler(serveStatics)).MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
        return strings.Contains(r.RequestURI, ".")
    })

    // home controller
    router.HandleFunc("/", useHandler(frontend.Home(useTemplate("home.html"))))

    // login controller
    router.HandleFunc("/login",          useHandler(frontend.Login(useTemplate("login.html"))))
    router.HandleFunc("/login/callback", useHandler(frontend.Callback(), mongo))

    // logout controller
    router.HandleFunc("/logout", useHandler(frontend.Logout()))

    // profile controller
    router.HandleFunc("/profile", useHandler(frontend.Profile(useTemplate("profile.html")), auth))

    // admin
    router.HandleFunc("/admin",       useHandler(backend.Home(useBackendTemplate("home.html")),   admin, mongo))
    router.HandleFunc("/admin/users", useHandler(backend.Users(useBackendTemplate("users.html")), admin, mongo))

    return router
}()

func serveStatics(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    file := vars["file"]

    http.ServeFile(w, r, "./public/" + file)
}