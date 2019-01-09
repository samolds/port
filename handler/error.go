// Copyright (C) 2018 Sam Olds

package handler

import (
	"net/http"

	"github.com/samolds/port/template"
)

func NotFound(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return template.Error.Render(w, "'"+r.URL.Path+"' cannot be found")
}

func UnsupportedMethod(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return template.Error.Render(w, "'"+r.Method+"' is unsupported")
}
