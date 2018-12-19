// Copyright (C) 2018 Sam Olds

package port

import (
	"log"
	"net/http"

	"github.com/samolds/port/handler"
	"github.com/samolds/port/httpError"
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
	mux := NewRichMux()
	mux.Handle("GET", "", httpError.Handler(handler.Home))
	mux.Handle("GET", "/now", httpError.Handler(handler.Now))
	mux.Handle("GET", "/links", httpError.Handler(handler.Links))
	//mux.Handle("GET", "/a/b/c", httpError.Handler(handler.Links))
	mux.HandleDir("GET", "/static", http.FileServer(http.Dir("./"+staticDir+"/")))

	log.Printf("%#v\n", mux.sub)
	log.Printf("%#v\n", mux.methods)

	//mux := http.NewServeMux()
	//mux.Handle("/", herr(handler.Home))
	//mux.Handle("/links", herr(handler.Links))
	//mux.Handle("/now", herr(handler.Now))
	//mux.Handle("/static/", http.FileServer(http.Dir("./"+staticDir+"/")))

	server := Server{Router: mux}
	return server, nil
}
