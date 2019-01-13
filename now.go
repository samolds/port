// Copyright (C) 2018 Sam Olds

package port

import (
	"context"
	stdtemplate "html/template"
	"net/http"

	"github.com/samolds/port/template"
)

func (s *Server) Now(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	now, err := s.db.GetMostRecentNowText(ctx)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["htmlText"] = stdtemplate.HTML(now.HTMLText)
	data["profileImgSrc"] = now.ProfileImgSrc
	return template.Now.Render(w, data)
}
