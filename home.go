// Copyright (C) 2018 Sam Olds

package port

import (
	"context"
	"net/http"

	"github.com/samolds/port/template"
)

func (_ *Server) Home(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	return template.Home.Render(w, nil)
}
