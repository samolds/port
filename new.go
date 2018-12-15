// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"
)

type Server struct {
	DB     string // TODO: unused
	Router http.Handler
}

var (
	staticDir = "src/github.com/samolds/port/static"
)

// New initializes a new http handler for this web server.
func NewServer() (Server, error) {
	// TODO: make a rich muxer that will return 404 on no match and double check
	//       method
	//mux := newNotFoundMux()
	mux := http.NewServeMux()
	mux.Handle("/", errHandler(home))
	mux.Handle("/links", errHandler(links))
	mux.Handle("/now", errHandler(now))
	mux.Handle("/static", http.FileServer(http.Dir("./"+staticDir+"/")))

	server := Server{Router: mux}
	return server, nil
}
