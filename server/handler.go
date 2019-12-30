// Copyright (C) 2018 - 2020 Sam Olds

package server

import (
	"context"
	stdtemplate "html/template"
	"net/http"

	"github.com/samolds/port/template"
)

func (_ *Server) Home(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	return template.Home.Render(w, nil)
}

func (s *Server) Link(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	links, err := s.db.GetAllLinks(ctx)
	if err != nil {
		return err
	}

	return template.Link.Render(w, links)
}

func (s *Server) Now(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	now, err := s.db.GetMostRecentNowText(ctx)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["htmlText"] = stdtemplate.HTML(now.HTMLText)
	data["profileImgSrc"] = now.ProfileImgSrc
	data["creationTime"] = now.CreationTime.Format("Jan _2, 2006")
	return template.Now.Render(w, data)
}

func (_ *Server) NotFound(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return template.Error.Render(w, "'"+r.URL.Path+"' cannot be found")
}

func (_ *Server) UnsupportedMethod(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return template.Error.Render(w, "'"+r.Method+"' is unsupported")
}
