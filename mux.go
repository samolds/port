// Copyright (C) 2018 Sam Olds

package port

import (
	"net/http"
	"strings"

	"github.com/samolds/port/httpError"
)

type RichMux struct {
	sub      map[string]*RichMux
	methods  map[string]http.Handler
	dirServe bool
}

func NewRichMux() *RichMux {
	return &RichMux{}
}

func (r *RichMux) handle(method string, pattern string, dirServe bool,
	handler http.Handler) {
	method = strings.ToUpper(method)
	pattern = strings.Trim(pattern, " /")
	parts := strings.Split(pattern, "/")

	if pattern == "" {
		// at a path leaf node, need to map a handler to a method
		if r.methods == nil {
			r.methods = make(map[string]http.Handler)
		}
		_, exists := r.methods[method]
		if exists {
			panic(method + " method is already registered here")
		}

		// the handler is registered to the method at the path
		r.methods[method] = handler
		r.dirServe = dirServe
		return
	}

	// need to traverse down or build the sub tree
	if r.sub == nil {
		r.sub = make(map[string]*RichMux)
	}

	// recurse
	subDir := parts[0]
	_, exists := r.sub[subDir]
	if !exists {
		// add the subdirectory iff it doesn't exist
		r.sub[subDir] = &RichMux{}
	}

	r.sub[subDir].handle(method, strings.Join(parts[1:], "/"), dirServe,
		handler)
}

func (r *RichMux) Handle(method string, pattern string, handler http.Handler) {
	r.handle(method, pattern, false, handler)
}

func (r *RichMux) HandleDir(method string, pattern string, handler http.Handler) {
	r.handle(method, pattern, true, handler)
}

func (m *RichMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpError.Handler(func(writer http.ResponseWriter, req *http.Request) error {
		path := strings.Trim(r.URL.Path, "/")

		ptr := m
		for path != "" {
			parts := strings.Split(path, "/")
			bit := parts[0]

			var exists bool
			ptr, exists = ptr.sub[bit]
			if !exists {
				return httpError.New(r.URL.Path+" cannot be found", http.StatusNotFound)
			}

			path = strings.Join(parts[1:], "/")
			if ptr.dirServe {
				// if this whole subdir should be served, stop stepping down the tree
				break
			}
		}

		h, exists := ptr.methods[r.Method]
		if !exists {
			return httpError.New(r.Method+" is unsupported",
				http.StatusMethodNotAllowed)
		}

		h.ServeHTTP(writer, req)
		return nil
	}).ServeHTTP(w, r)

}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

/*
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
*/
