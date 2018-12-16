// Copyright (C) 2018 Sam Olds

package port

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/samolds/port/httpError"
)

// TODO: make notFoundMux a richMux
// TODO: this doesn't work for static file serving!
// not found mux
// vvvvvvvvvvvvvvvvv
type notFoundMux struct {
	*http.ServeMux
	routes map[string]interface{}
}

func newNotFoundMux() *notFoundMux {
	mux := http.NewServeMux()
	routes := make(map[string]interface{})
	return &notFoundMux{ServeMux: mux, routes: routes}
}

func (h *notFoundMux) Handle(pattern string, handler http.Handler) {
	_, exists := h.routes[pattern]
	if exists {
		panic("duplicate routes registered")
	}

	h.routes[pattern] = new(interface{})
	h.ServeMux.Handle(pattern, handler)
}

func (h *notFoundMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, exists := h.routes[r.URL.Path]
	if !exists {
		log.Printf("error: 404 not found: %s", r.URL.Path)
		http.Error(w, fmt.Sprintf("404 path %s not found", r.URL.Path),
			http.StatusNotFound)
		return
	}

	h.ServeMux.ServeHTTP(w, r)
}

// ^^^^^^^^^^^^^^^^^
// catchPanic can be used to catch panics and turn them into errors.
func catchPanic(errRef *error) {
	r := recover()
	if r == nil {
		return
	}
	err, ok := r.(error)
	if ok {
		*errRef = errors.New(fmt.Sprintf("panic: %v", err))
		return
	}
	*errRef = errors.New(fmt.Sprintf("%v", r))
}

// ^^^^^^^^^^^^^^^^^

// err handler
// vvvvvvvvvvvvvvvvv

// errHandler is a handler with a returned error.
type errHandler func(w http.ResponseWriter, req *http.Request) error

// ServeHTTP is so that errHandler is a proper handler and can be used with
// incoming requests.
func (eh errHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.URL.Path = %s", r.URL.Path)
	err := func() (err error) {
		defer catchPanic(&err)
		return eh(w, r)
	}()

	herr := httpError.Catch(err)
	if err != nil {
		log.Printf("error: %s", err)
		http.Error(w, herr.Error(), herr.StatusCode)
	}
}
