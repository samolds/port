// Copyright (C) 2018 Sam Olds

package httperror

import (
	"context"
	"log"
	"net/http"
)

// Handler is a handler with a returned error.
type Handler func(context.Context, http.ResponseWriter, *http.Request) error

// ServeHTTP is so that Handler is a proper handler and can be used with
// incoming requests.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := func() (err error) {
		defer CatchPanic(&err)
		return h(context.Background(), w, r)
	}()

	herr := Catch(err)
	if err != nil {
		log.Printf("error: %s", err)
		w.WriteHeader(herr.StatusCode)
		http.Error(w, herr.Error(), herr.StatusCode)
	}
}
