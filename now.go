// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"

	"github.com/samolds/port/template"
)

func now(w http.ResponseWriter, r *http.Request) error {
	err := template.Render(w, r, template.Now, map[string]interface{}{})
	if err != nil {
		return err
	}
	return nil
}
