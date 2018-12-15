// Copyright (C) 2018 Sam Olds

package port

import (
	"fmt"
	"log"
	"net/http"
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
