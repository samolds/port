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
)

// main you know what did is.
func main() {
	err := runner()
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
}

// runner is the "real" main so that we can be idiomatic and just return errors
// everywhere, and main is just responsible for calling log.Fatalf on errors.
func runner() error {
	server, err := port.NewServer()
	if err != nil {
		return err
	}

	return http.ListenAndServe(*tcpPort, httplog.LogResponses(server.Router))
}
