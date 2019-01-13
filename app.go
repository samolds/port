// Copyright (C) 2018 Sam Olds

package port

import (
	"context"
	"log"
	"net/http"

	"github.com/samolds/port/database"
	"github.com/samolds/port/httperror"
	"github.com/samolds/port/httpmux"
	"github.com/samolds/port/template"
)

type Options struct {
	StaticDir    string
	GAEProjectID string
	GAECredFile  string
}

type Server struct {
	router http.Handler
	db     *database.DB
}

// New initializes a new http handler for this web server.
func NewServer(ctx context.Context, opts Options) (*Server, error) {
	server := &Server{}

	db, err := database.New(ctx, opts.GAEProjectID, opts.GAECredFile)
	if err != nil {
		return nil, err
	}
	server.db = db

	mux := httpmux.New()
	mux.RegisterNotFoundHandler(herr(server.NotFound))
	mux.RegisterUnsupportedMethodHandler(herr(server.UnsupportedMethod))

	mux.Handle("GET", "/", herr(server.Home))
	mux.Handle("GET", "/now", herr(server.Now))
	mux.Handle("GET", "/links", herr(server.Link))
	server.router = mux

	if opts.StaticDir != "" {
		// optional because a separate static file server might be used
		mux.HandleDir("GET", "/static", http.FileServer(http.Dir(opts.StaticDir)))
	}

	return server, nil
}

// herr wraps all handlers as an httperror.Handler and will attempt to catch
// any error and render them in a nice error template. worse case, it will
// display the raw error not as a template
func herr(h func(context.Context, http.ResponseWriter, *http.Request) error) (
	_ httperror.Handler) {
	return httperror.Handler(func(ctx context.Context, w http.ResponseWriter,
		r *http.Request) error {
		err := h(ctx, w, r)
		if err != nil {
			log.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			subErr := template.Error.Render(w, "internal server error: "+err.Error())
			if subErr != nil {
				log.Println(subErr.Error())
				return err
			}
		}
		return nil
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
