// Copyright (C) 2018 Sam Olds

package template

import (
	"bytes"
	stdtemplate "html/template"
	"io"
	"net/http"
)

const (
	// make sure to update the definition in template/base.html if this changes
	baseTmplName = "base"

	homeTmplName  = "home"
	linksTmplName = "links"
	nowTmplName   = "now"
)

var (
	// TODO: make this not hardcoded
	templateDir = "src/github.com/samolds/port/template/pages/"

	baseTmplFile  = templateDir + baseTmplName + ".html"
	homeTmplFile  = templateDir + homeTmplName + ".html"
	linksTmplFile = templateDir + linksTmplName + ".html"
	nowTmplFile   = templateDir + nowTmplName + ".html"

	// the exported templates that are available to render
	Home  = mustParse(homeTmplFile)
	Links = mustParse(linksTmplFile)
	Now   = mustParse(nowTmplFile)
)

type tmpl struct {
	*stdtemplate.Template
}

func mustParse(tmplFile string) tmpl {
	t := stdtemplate.Must(stdtemplate.New(baseTmplName).ParseFiles(baseTmplFile,
		tmplFile))
	return tmpl{Template: t}
}

// Render writes the template out to the response writer (or any errors that
// come up), with values as the template value. If there were errors rendering
// the template, they will not be written out to the writer.
func (t tmpl) Render(w http.ResponseWriter, values interface{}) error {
	var buf bytes.Buffer
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK) // TODO: currently redundant
	err := t.Template.Execute(&buf, values)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, &buf)
	return err
}
