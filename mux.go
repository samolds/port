// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"
	"strings"
	"sync"

	"github.com/samolds/port/httpError"
)

type richMux struct {
	mu     sync.Mutex
	routes map[string]muxEntry
}

type muxEntry struct {
	http.Handler

	isDir  bool
	path   string
	method string
}

func NewRichMux() richMux {
	return richMux{routes: make(map[string]muxEntry)}
}

func (m richMux) handle(method string, path string, isDir bool,
	handler http.Handler) {
	m.mu.Lock()
	defer m.mu.Unlock()

	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" && method != "PUT" &&
		method != "DELETE" {
		panic(method + " is an unsupported method")
	}

	path = strings.Trim(path, "/")
	_, exists := m.routes[path]
	if exists {
		panic(path + " path already registered")
	}

	m.routes[path] = muxEntry{
		Handler: handler,
		isDir:   isDir,
		path:    path,
		method:  method,
	}
}

func (m richMux) Handle(method string, path string,
	handler http.Handler) {
	m.handle(method, path, false, handler)
}

func (m richMux) HandleDir(method string, path string,
	handler http.Handler) {
	m.handle(method, path, true, handler)
}

// TODO: finish this
func (m richMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpError.Handler(func(writer http.ResponseWriter, req *http.Request) error {
		path := strings.Trim(r.URL.Path, "/")

		entry, exists := m.routes[path]
		if !exists {
			return httpError.New(r.URL.Path+" cannot be found", http.StatusNotFound)
		}

		if entry.method != r.Method {
			return httpError.New(r.Method+" is unsupported",
				http.StatusMethodNotAllowed)
		}

		if entry.isDir {
			entry.Handler.ServeHTTP(writer, req)
			return nil
		}

		entry.Handler.ServeHTTP(writer, req)
		return nil
	}).ServeHTTP(w, r)
}
