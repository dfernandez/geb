package decorator

import (
	"net/http"
	"log"
	"time"
)

type Logger struct {
	Logger *log.Logger
}

func (l Logger) Do(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		defer func() {
			t1 := time.Now()
			l.Logger.Printf("%s Took: %s", r.RequestURI, t1.Sub(t0).String())
		}()

		h(w, r)
	}
}
