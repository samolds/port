// Copyright (C) 2018 Sam Olds

package port

import (
	"errors"
	"fmt"
	//"gopkg.in/webhelp.v1/whmux"
	"log"
	"net/http"
)

type Server struct {
	DB     string // TODO: unused
	Router http.Handler
}

type httpError struct {
	err        error
	statusCode int
}

func (herr httpError) status() string {
	// TODO: get http status based on statusCode
	return "TODO"
}

var (
	staticDir = "src/github.com/samolds/port/static"
)

// err handler
// vvvvvvvvvvvvvvvvv

// errHandler is a handler with a returned error.
// TODO: return an httpError instead of a standard error?
type errHandler func(w http.ResponseWriter, req *http.Request) error

// ServeHTTP is so that errHandler is a proper handler and can be used with
// incoming requests.
func (eh errHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.URL.Path = %s", r.URL.Path)
	err := func() (err error) {
		defer catchPanic(&err)
		return eh(w, r)
	}()

	if err != nil {
		// TODO: add error classes so we know what kind of error should be thrown
		//       here
		log.Printf("error: %s", err)
		http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
	}
}

// ^^^^^^^^^^^^^^^^^

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
	//_, exists := h.routes[r.URL.Path]
	//if !exists {
	//	log.Printf("error: 404 not found: %s", r.URL.Path)
	//	http.Error(w, fmt.Sprintf("404 path %s not found", r.URL.Path),
	//		http.StatusNotFound)
	//	return
	//}

	h.ServeMux.ServeHTTP(w, r)
}

// ^^^^^^^^^^^^^^^^^

// New initializes a new http handler for this web server.
func NewServer() (Server, error) {
	mux := newNotFoundMux()
	mux.Handle("/", errHandler(home))
	mux.Handle("/links", errHandler(links))
	mux.Handle("/now", errHandler(now))
	mux.Handle("/static", http.FileServer(http.Dir("./"+staticDir+"/")))

	server := Server{Router: mux}
	return server, nil
}

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
