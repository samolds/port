// Copyright (C) 2018 Sam Olds

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/samolds/port"
	"github.com/samolds/port/httplog"
)

var (
	tcpPort = flag.String("port", ":8080", "port to listen on")

	// like:
	//   /Users/samolds/projects/go/src/github.com/samolds/port/static
	staticDir = flag.String("static-dir", "",
		"absolute path to directory with static assets")
)

// main you know what did is.
func main() {
	flag.Parse()
	err := runner()
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
}

// runner is the "real" main so that we can be idiomatic and just return errors
// everywhere, and main is just responsible for calling log.Fatalf on errors.
func runner() error {
	opts := port.Options{
		StaticDir: *staticDir,
	}

	server, err := port.NewServer(opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(*tcpPort, httplog.LogResponses(server.Router))
}
