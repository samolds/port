// Copyright (C) 2018 Sam Olds

package wwwrun

import (
  "net/http"
  "flag"

  "github.com/samolds/www"
)

var (
	port = flag.String("port", ":8080", "port to listen on")
)

func main() {
  err := runner()
  if err != nil {
    Fatalf("failed: %v", err)
  }
}

func runner() error {
  handlers, err := www.New()
  if err != nil {
    return err
  }

  http.ListenAndServe(*port, handlers)
  return nil
}
