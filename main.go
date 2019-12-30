// Copyright (C) 2018 - 2020 Sam Olds

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

// TODO: support a config file
var (
	tcpPort = flag.String("port", "", "port to listen on")

	staticDir = flag.String("static-dir", "",
		"absolute path to directory with static assets")

	gaeProjectID = flag.String("gae-project-id", "",
		"the project id for this google app engine project")

	gaeCredFile = flag.String("gae-cred-file", "",
		"the path to the credential file for google app engine")

	relHTMLTmplDir = flag.String("rel-html-tmpl-dir", "",
		"the relative path to the directory with all of the html templates")
)

// GAE uses environment variables
const (
	tcpPortEV        = "PORT"
	staticDirEV      = "STATIC_DIR"
	gaeProjectIDEV   = "GAE_PROJECT_ID"
	gaeCredFileEV    = "GAE_CRED_FILE"
	relHTMLTmplDirEV = "REL_HTML_TMPL_DIR"
)

// main you know what this is.
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
		StaticDir:      *staticDir,
		GAEProjectID:   *gaeProjectID,
		GAECredFile:    *gaeCredFile,
		RelHTMLTmplDir: *relHTMLTmplDir,
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
	flags := *tcpPort != "" ||
		*staticDir != "" ||
		*gaeProjectID != "" ||
		*gaeCredFile != "" ||
		*relHTMLTmplDir != ""
	if flags {
		return
	}

	*tcpPort = os.Getenv(tcpPortEV)
	*staticDir = os.Getenv(staticDirEV)
	*gaeProjectID = os.Getenv(gaeProjectIDEV)
	*gaeCredFile = os.Getenv(gaeCredFileEV)
	*relHTMLTmplDir = os.Getenv(relHTMLTmplDirEV)

	// prepend ":" to the port if it wasn't provided by dumb gae
	p := *tcpPort
	if len(p) == 0 {
		panic("port is required")
	}

	if !strings.HasPrefix(p, ":") {
		*tcpPort = ":" + p
	}
}
