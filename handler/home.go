// Copyright (C) 2018 Sam Olds

package handler

import (
	"net/http"

	"github.com/samolds/port/template"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	err := template.Home.Render(w, map[string]interface{}{})
	if err != nil {
		return err
	}
	return nil
}
