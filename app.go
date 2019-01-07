// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"

	"github.com/samolds/port/handler"
	"github.com/samolds/port/httperror"
	"github.com/samolds/port/httpmux"
)

type Server struct {
	DB     string // TODO: unused
	Router http.Handler
}

var (
	// TODO: make this not hardcoded
	staticDir = "src/github.com/samolds/port/static"
)

// New initializes a new http handler for this web server.
func NewServer() (Server, error) {
	mux := httpmux.New()
	mux.RegisterNotFoundHandler(herr(handler.NotFound))
	mux.RegisterUnsupportedMethodHandler(herr(handler.UnsupportedMethod))

	mux.Handle("GET", "", herr(handler.Home))
	mux.Handle("GET", "/now", herr(handler.Now))
	mux.Handle("GET", "/links", herr(handler.Links))
	mux.HandleDir("GET", "/static", http.FileServer(http.Dir("./"+staticDir+"/")))

	server := Server{Router: mux}
	return server, nil
}

func herr(h func(http.ResponseWriter, *http.Request) error) httperror.Handler {
	return httperror.Handler(h)
}
