// Copyright (C) 2018 - 2020 Sam Olds

package httperror

import (
	"errors"
	"fmt"
	"net/http"
)

// Error is a rich error with a field for an http status code
type Error struct {
	Err        string
	StatusCode int
}

// Error will return a string representation of the rich Error
func (err *Error) Error() string {
	return fmt.Sprintf("%d %s: {%s}", err.StatusCode,
		http.StatusText(err.StatusCode), err.Err)
}

// New returns a new rich error
func New(err string, statusCode int) *Error {
	return &Error{Err: err, StatusCode: statusCode}
}

// Catch will detect a rich Error from and error type, or wrap the error as a
// rich Error
func Catch(err error) *Error {
	if err == nil {
		return nil
	}

	herr, ok := err.(*Error)
	if !ok {
		return New(err.Error(), http.StatusInternalServerError)
	}
	return herr
}

// CatchPanic can be used to catch panics and turn them into errors.
func CatchPanic(errRef *error) {
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
