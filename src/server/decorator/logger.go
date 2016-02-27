package decorator

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"time"
)

type Logger struct {
	Logger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{log.New()}
}

func (l Logger) Do(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		h(w, r)
		l.Logger.Info(fmt.Sprintf("%s %s %s %s", r.RemoteAddr, r.Method, r.RequestURI, time.Now().Sub(t0).String()))
	}
}
