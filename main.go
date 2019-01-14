// Copyright (C) 2018 - 2019 Sam Olds

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/samolds/port/httplog"
	"github.com/samolds/port/server"
)

var (
	tcpPort = flag.String("port", "", "port to listen on")

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
	envVarCheck()
	err := runner()
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
}

// runner is the "real" main so that we can be idiomatic and just return errors
// everywhere, and main is just responsible for calling log.Fatalf on errors.
func runner() error {
	opts := server.Options{
		StaticDir:    *staticDir,
		GAEProjectID: *gaeProjectID,
		GAECredFile:  *gaeCredFile,
	}

	ctx := context.Background()
	server, err := server.New(ctx, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(*tcpPort, httplog.LogResponses(server))
}

// envVarCheck is to be called after flag.Parse() to check to see if no flags
// were provided. if no flags were provided, it's probably because it's running
// on stupid GAE and is expecting environment variables instead. dumb
func envVarCheck() {
	flags := *tcpPort != "" || *staticDir != "" || *gaeProjectID != "" ||
		*gaeCredFile != ""
	if flags {
		return
	}

	*tcpPort = os.Getenv("PORT")
	*staticDir = os.Getenv("STATIC_DIR")
	*gaeProjectID = os.Getenv("GAE_PROJECT_ID")
	*gaeCredFile = os.Getenv("GAE_CRED_FILE")

	// prepend ":" to the port if it wasn't provided by dumb gae
	p := *tcpPort
	if len(p) == 0 {
		panic("port is required")
	}

	if !strings.HasPrefix(p, ":") {
		*tcpPort = ":" + p
	}
}
