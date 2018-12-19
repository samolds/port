// Copyright (C) 2018 Sam Olds

package handler

import (
	"net/http"

	"github.com/samolds/port/httpError"
	"github.com/samolds/port/template"
)

func Links(w http.ResponseWriter, r *http.Request) error {
	err := template.Links.Render(w, map[string]interface{}{})
	if err != nil {
		return httpError.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
