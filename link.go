// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"

	"github.com/samolds/port/template"
)

func links(w http.ResponseWriter, r *http.Request) error {
	err := template.Links.Render(w, map[string]interface{}{})
	if err != nil {
		return err
	}
	return nil
}
