// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"

	"github.com/samolds/port/template"
)

func home(w http.ResponseWriter, r *http.Request) error {
	err := template.Render(w, r, template.Home, map[string]interface{}{})
	if err != nil {
		return err
	}
	return nil
}
