// Copyright (C) 2018 Sam Olds

package port

import (
	"log"
	"net/http"

	"github.com/samolds/port/handler"
	"github.com/samolds/port/httperror"
	"github.com/samolds/port/httpmux"
	"github.com/samolds/port/template"
)

type Options struct {
	StaticDir string
}

type Server struct {
	DB     string // TODO: unused
	Router http.Handler
}

// New initializes a new http handler for this web server.
func NewServer(opts Options) (Server, error) {
	mux := httpmux.New()
	mux.RegisterNotFoundHandler(herr(handler.NotFound))
	mux.RegisterUnsupportedMethodHandler(herr(handler.UnsupportedMethod))

	mux.Handle("GET", "/", herr(handler.Home))
	mux.Handle("GET", "/now", herr(handler.Now))
	mux.Handle("GET", "/links", herr(handler.Link))

	if opts.StaticDir != "" {
		mux.HandleDir("GET", "/static", http.FileServer(http.Dir(opts.StaticDir)))
	}

	server := Server{Router: mux}
	return server, nil
}

// herr wraps all handlers as an httperror.Handler and will attempt to catch
// any error and render them in a nice error template. worse case, it will
// display the raw error not as a template
func herr(h func(http.ResponseWriter, *http.Request) error) httperror.Handler {
	return httperror.Handler(func(w http.ResponseWriter, r *http.Request) error {
		err := h(w, r)
		if err != nil {
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
