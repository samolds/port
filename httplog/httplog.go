// Copyright (C) 2018 - 2020 Sam Olds

package httplog

import (
	"log"
	"net/http"
	"time"
)

// responseWriterWrapper is used to keep track of the status code written back
// for logging purposes
type responseWriterWrapper struct {
	rw            http.ResponseWriter
	statusCode    int
	wroteHeader   bool
	contentLength int
}

func (w *responseWriterWrapper) Header() http.Header {
	return w.rw.Header()
}

func (w *responseWriterWrapper) Write(data []byte) (_ int, err error) {
	w.contentLength, err = w.rw.Write(data)
	return w.contentLength, err
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	if w.wroteHeader {
		return
	}

	w.statusCode = statusCode
	w.rw.WriteHeader(statusCode)
	w.wroteHeader = true
}

// LogResponses takes a Handler and makes it log responses
func LogResponses(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &responseWriterWrapper{
			rw:         w,
			statusCode: http.StatusOK,
		}

		start := time.Now()
		h.ServeHTTP(ww, r)
		log.Printf(`%s %d %#v %d %d`, r.Method, ww.statusCode, r.RequestURI,
			ww.contentLength, time.Since(start))
	})
}
