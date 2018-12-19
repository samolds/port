// Copyright (C) 2018 Sam Olds

package httpError

import (
	"log"
	"net/http"

	"github.com/samolds/port/template"
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
		w.WriteHeader(herr.StatusCode)
		err = template.Error.Render(w, herr.Error())
		if err != nil {
			http.Error(w, err.Error(), herr.StatusCode)
		}
	}
}
