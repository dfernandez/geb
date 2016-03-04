package server

import (
	"net/http"
	"github.com/dfernandez/geb/src/server/decorator"
	"sort"
)

type httpHandlerDecorator interface {
	Do(http.HandlerFunc) http.HandlerFunc
}

func useHandler(h http.HandlerFunc, decors ...httpHandlerDecorator) http.HandlerFunc {
	context := decorator.NewContext()
	logger  := decorator.NewLogger()
	recover := decorator.NewRecover()

	n        := len(decors)
	handlers := make(map[int]httpHandlerDecorator, (n + 3))

	var userKeys []int
	for k := range decors {
		userKeys = append(userKeys, k)
	}

	sort.Ints(userKeys)

	for _, k := range userKeys {
		handlers[k] = decors[k]
	}

	handlers[(n+1)] = context
	handlers[(n+2)] = logger
	handlers[(n+3)] = recover

	var keys []int
	for k := range handlers {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		h = handlers[k].Do(h)
	}

	return h
}