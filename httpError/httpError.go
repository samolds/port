// Copyright (C) 2018 Sam Olds

package httpError

import (
	"fmt"
	"net/http"
)

type Error struct {
	Err        error
	Msg        string
	StatusCode int
}

func (err *Error) Error() string {
	return fmt.Sprintf("%d %s: {%s} %s", err.StatusCode,
		http.StatusText(err.StatusCode), err.Err.Error(), err.Msg)
}

func New(err error, msg string, statusCode int) *Error {
	return &Error{Err: err, Msg: msg, StatusCode: statusCode}
}

func Catch(err error) *Error {
	herr, ok := err.(*Error)
	if !ok {
		return New(err, "", http.StatusInternalServerError)
	}
	return herr
}
