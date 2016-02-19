package controller

import (
	"net/http"
	"fmt"
)

func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "Internal server error")
}
