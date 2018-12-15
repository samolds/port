// Copyright (C) 2018 Sam Olds

package port

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// httpError
// vvvvvvvvvvvvvvvvv

type httpError struct {
	err        error
	msg        string
	statusCode int
}

func (herr *httpError) Error() string {
	return fmt.Sprintf("%d %s: {%s} %s", herr.statusCode,
		http.StatusText(herr.statusCode), herr.err.Error(), herr.msg)
}

func newHTTPError(err error, msg string, statusCode int) *httpError {
	return &httpError{err: err, msg: msg, statusCode: statusCode}
}

func catchHTTPError(err error) *httpError {
	herr, ok := err.(*httpError)
	if !ok {
		return newHTTPError(err, "", http.StatusInternalServerError)
	}
	return herr

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

// ^^^^^^^^^^^^^^^^^

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

	herr := catchHTTPError(err)
	if err != nil {
		log.Printf("error: %s", err)
		http.Error(w, herr.Error(), herr.statusCode)
	}
}
