// Copyright (C) 2018 Sam Olds

package httpmux

import (
	"net/http"
	"strings"
)

type SimpleMux struct {
	mux
	notFoundHandler          http.Handler
	unsupportedMethodHandler http.Handler
}

type mux struct {
	sub      map[string]*mux
	methods  map[string]http.Handler
	dirServe bool
}

func New() *SimpleMux {
	return &SimpleMux{}
}

// handle will store the provided method and pattern in the muxer, building a
// tree that can be used by ServeHTTP to handle the requests
//
// NOTE: this simple muxer does not currently support url parameters
func (m *mux) handle(method string, pattern string, dirServe bool,
	handler http.Handler) {
	method = strings.ToUpper(method)
	pattern = strings.Trim(pattern, " /")
	parts := strings.Split(pattern, "/")

	if pattern == "" {
		// at a path leaf node, need to map a handler to a method
		if m.methods == nil {
			m.methods = make(map[string]http.Handler)
		}

		_, exists := m.methods[method]
		if exists {
			panic(method + " method is already registered here")
		}

		// the handler is registered to the method at the path
		m.methods[method] = handler
		m.dirServe = dirServe
		return
	}

	// need to traverse down or build the sub tree
	if m.sub == nil {
		m.sub = make(map[string]*mux)
	}

	// recurse
	subDir := parts[0]
	_, exists := m.sub[subDir]
	if !exists {
		// add the subdirectory iff it doesn't exist
		m.sub[subDir] = &mux{}
	}

	m.sub[subDir].handle(method, strings.Join(parts[1:], "/"), dirServe, handler)
}

func (m *SimpleMux) Handle(me string, p string, h http.Handler) {
	m.mux.handle(me, p, false, h)
}

func (m *SimpleMux) HandleDir(me string, p string, h http.Handler) {
	h = http.StripPrefix(p, h)
	m.mux.handle(me, p, true, h)
}

func (m *SimpleMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")

	ptr := &m.mux
	for path != "" {
		parts := strings.Split(path, "/")
		bit := parts[0]

		var exists bool
		ptr, exists = ptr.sub[bit]
		if !exists {
			if m.notFoundHandler == nil {
				http.Error(w, "'"+r.URL.Path+"' cannot be found", http.StatusNotFound)
			} else {
				m.notFoundHandler.ServeHTTP(w, r)
			}
			return
		}

		path = strings.Join(parts[1:], "/")
		if ptr.dirServe {
			// if this whole subdir should be served, stop stepping down the tree
			break
		}
	}

	h, exists := ptr.methods[r.Method]
	if !exists {
		if m.unsupportedMethodHandler == nil {
			http.Error(w, "'"+r.Method+"' is unsupported",
				http.StatusMethodNotAllowed)
		} else {
			m.unsupportedMethodHandler.ServeHTTP(w, r)
		}
		return
	}

	// TODO: somehow catch if the file server throws an error and catch it here?
	if ptr.dirServe {
	}

	h.ServeHTTP(w, r)
}

func (m *SimpleMux) RegisterNotFoundHandler(h http.Handler) {
	m.notFoundHandler = h
}

func (m *SimpleMux) RegisterUnsupportedMethodHandler(h http.Handler) {
	m.unsupportedMethodHandler = h
}
