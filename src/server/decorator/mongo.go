package decorator

import (
	"net/http"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"github.com/dfernandez/geb/config"
)

type Mongo struct {
}

func NewMongo() *Mongo {
	return &Mongo{}
}

func (m Mongo) Do(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := mgo.Dial(config.MongoServer)
		if err != nil {
			panic(err)
		}

		session.SetMode(mgo.Monotonic, true)

		context.Set(r, "mongoDB", session)
		h(w, r)

		session.Close()
	}
}
