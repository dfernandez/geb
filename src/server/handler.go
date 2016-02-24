package server

import (
	"net/http"
)

type httpHandlerDecorator interface {
	Do(http.HandlerFunc) http.HandlerFunc
}

func useHandler(h http.HandlerFunc, decors ...httpHandlerDecorator) http.HandlerFunc {
	for _, decorator := range decors {
		h = decorator.Do(h)
	}

	return h
}
