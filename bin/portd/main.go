// Copyright (C) 2018 Sam Olds

package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/samolds/port"
	"github.com/samolds/port/httplog"
)

var (
	tcpPort = flag.String("port", ":8080", "port to listen on")

	staticDir = flag.String("static-dir", "",
		"absolute path to directory with static assets")

	gaeProjectID = flag.String("gae-project-id", "",
		"the project id for this google app engine project")

	gaeCredFile = flag.String("gae-cred-file", "",
		"the path to the credential file for google app engine")
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
		StaticDir:    *staticDir,
		GAEProjectID: *gaeProjectID,
		GAECredFile:  *gaeCredFile,
	}

	ctx := context.Background()
	server, err := port.NewServer(ctx, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(*tcpPort, httplog.LogResponses(server))
}
