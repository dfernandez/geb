package controllers

import (
	"net/http"
	"fmt"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "404")
}
