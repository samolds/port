// Copyright (C) 2018 Sam Olds

package httpError

import (
	"log"
	"net/http"
)

// Handler is a handler with a returned error.
type Handler func(w http.ResponseWriter, req *http.Request) error

// ServeHTTP is so that Handler is a proper handler and can be used with
// incoming requests.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.URL.Path = %s", r.URL.Path)
	err := func() (err error) {
		defer CatchPanic(&err)
		return h(w, r)
	}()

	herr := Catch(err)
	if err != nil {
		log.Printf("error: %s", err)
		http.Error(w, herr.Error(), herr.StatusCode)
	}
}
