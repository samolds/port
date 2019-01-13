# Port
Source code for my personal web server

Too many analogies with port - like a spaceport or portal to the web for Sam's
portfolio.

### About
This is an exercise in writing a simple, clean, idiomatic web server with only
the std lib.

### Setup
```sh
go get github.com/samolds/port
```

### Run Dev Version
```sh
cd $GOPATH
go install .../portd && ./bin/portd


go install .../portd && ./bin/portd
  --port=":8080"
  --static-dir="/Users/samolds/projects/go/src/github.com/samolds/port/bin/portd/static"
  --gae-project-id="samolds"
  --gae-cred-file="/Users/samolds/projects/go/src/github.com/samolds/port/assetdump/gae_cred_file_samolds.json"
```

### Interesting Libraries to explore
* github.com/samolds/port/template
* github.com/GoogleCloudPlatform/golang-samples/appengine/go11x/static/
* gopkg.in/webhelp.v1/whmux
* github.com/spacemonkeygo/spacelog
* gopkg.in/spacemonkeygo/monkit.v2
* github.com/zeebo/errs
* github.com/go-chi/chi
* cloud.google.com/go
* github.com/GoogleCloudPlatform/golang-samples/getting-started/bookshelf/app
