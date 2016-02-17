package decorator

import "net/http"

type HttpHandlerDecorator interface {
	Do(http.HandlerFunc) http.HandlerFunc
}

func Decorate(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for _, decorator := range decors {
		h = decorator.Do(h)
	}
	return h
}