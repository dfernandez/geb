package decorator

import (
	"net/http"
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	Logger *log.Logger
}

func NewLogger() *Logger {
	file, _ := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	multi := io.MultiWriter(file, os.Stdout)

	return &Logger{log.New(multi, "INFO: ", log.LstdFlags)}
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
