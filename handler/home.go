// Copyright (C) 2018 Sam Olds

package handler

import (
	"net/http"

	"github.com/samolds/port/template"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	return template.Home.Render(w, nil)
}
