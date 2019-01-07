// Copyright (C) 2018 Sam Olds

package httpmux

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samolds/port/testhelp"
)

func initHandler(tt *testhelp.T, responseBody string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, responseBody)
		tt.AssertNoError(err)
	})
}

func run(tt *testhelp.T, handler http.Handler, method string, path string,
	statusCode int, expResp string) {
	// create an empty request with the provided method to the provided path
	req, err := http.NewRequest(method, path, nil)
	tt.AssertNoError(err)

	// record the response while running the request throught the handler
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	// make sure it routed correctly
	tt.AssertEqual(rec.Code, statusCode)
	if statusCode == http.StatusOK {
		tt.AssertEqual(expResp, rec.Body.String())
	}
}

// TestMethodHome tests the muxer for two different methods on the same route
func TestMethodHome(t *testing.T) {
	tt := testhelp.New(t)
	getHandler := initHandler(tt, "GET home")
	postHandler := initHandler(tt, "POST home")

	// initialize the mux with two handlers
	mux := New()
	mux.Handle("GET", "/", getHandler)
	mux.Handle("POST", "/", postHandler)

	run(tt, mux, "GET", "/", http.StatusOK, "GET home")
	run(tt, mux, "POST", "/", http.StatusOK, "POST home")
}

// TestNested tests multiple handlers on nested routes
func TestNested(t *testing.T) {
	tt := testhelp.New(t)
	homeHandler := initHandler(tt, "GET home")
	subHandler := initHandler(tt, "GET sub")
	subSubHandler := initHandler(tt, "GET sub sub")

	// initialize the mux with three handlers
	mux := New()
	mux.Handle("GET", "/", homeHandler)
	mux.Handle("GET", "/sub", subHandler)
	mux.Handle("GET", "/sub/sub", subSubHandler)

	run(tt, mux, "GET", "", http.StatusOK, "GET home")
	run(tt, mux, "GET", "/sub", http.StatusOK, "GET sub")
	run(tt, mux, "GET", "/sub/sub", http.StatusOK, "GET sub sub")
}

// TestNotFound tests not finding a route that doesn't exists
func TestNotFound(t *testing.T) {
	tt := testhelp.New(t)
	homeHandler := initHandler(tt, "GET home")

	// initialize the mux with three handlers
	mux := New()
	mux.Handle("GET", "/", homeHandler)

	run(tt, mux, "GET", "/notRealPath", http.StatusNotFound, "")
}

// TestFlatHandlers tests multiple handlers all at the same depth
func TestFlatHandlers(t *testing.T) {
	tt := testhelp.New(t)
	oneHandler := initHandler(tt, "GET one")
	twoHandler := initHandler(tt, "GET two")
	threeHandler := initHandler(tt, "GET three")

	// initialize the mux with three handlers
	mux := New()
	mux.Handle("GET", "/one/", oneHandler)
	mux.Handle("GET", "two/", twoHandler)
	mux.Handle("GET", "three", threeHandler)

	run(tt, mux, "GET", "one", http.StatusOK, "GET one")
	run(tt, mux, "GET", "/two", http.StatusOK, "GET two")
	run(tt, mux, "GET", "/three/", http.StatusOK, "GET three")
}
