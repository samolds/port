// Copyright (C) 2018 Sam Olds

package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/samolds/port"
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

	return http.ListenAndServe(*tcpPort, logResponses(server.Router))
}

// logResponses takes a Handler and makes it log responses
func logResponses(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &responseWriterWrapper{rw: w}
		start := time.Now()
		h.ServeHTTP(ww, r)
		log.Printf(`%s %d %#v %d %d`, r.Method, ww.statusCode, r.RequestURI,
			r.ContentLength, time.Since(start))
	})
}

// responseWriterWrapper is used to keep track of the status code written back
// for logging purposes
type responseWriterWrapper struct {
	rw         http.ResponseWriter
	statusCode int
}

func (w *responseWriterWrapper) Header() http.Header {
	return w.rw.Header()
}

func (w *responseWriterWrapper) Write(data []byte) (int, error) {
	return w.rw.Write(data)
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.rw.WriteHeader(statusCode)
}
