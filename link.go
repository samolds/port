// Copyright (C) 2018 Sam Olds

package port

import (
	"context"
	"net/http"

	"github.com/samolds/port/template"
)

func (s *Server) Link(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	links, err := s.db.GetAllLinks(ctx)
	if err != nil {
		return err
	}

	return template.Link.Render(w, links)
}
