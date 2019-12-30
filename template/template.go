// Copyright (C) 2018 - 2020 Sam Olds

package template

import (
	"bytes"
	"errors"
	stdtemplate "html/template"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

const (
	// make sure to update the definition in template/base.html if this changes
	baseTmplName = "base"
)

var (
	baseTmplFile  = baseTmplName + ".html"
	homeTmplFile  = "home.html"
	linkTmplFile  = "link.html"
	nowTmplFile   = "now.html"
	errorTmplFile = "error.html"

	// the exported templates that are available to render
	Home  tmpl
	Link  tmpl
	Now   tmpl
	Error tmpl
)

func Initialize(relTmplDir string) error {
	if relTmplDir == "" {
		return errors.New("a path to the directory of templates is required")
	}

	templateDir, err := filepath.Abs(relTmplDir)
	if err != nil {
		return err
	}

	Home = mustParse(templateDir, homeTmplFile)
	Link = mustParse(templateDir, linkTmplFile)
	Now = mustParse(templateDir, nowTmplFile)
	Error = mustParse(templateDir, errorTmplFile)
	return nil
}

type tmpl struct {
	*stdtemplate.Template
}

func mustParse(templateDir string, tmplFile string) tmpl {
	base := path.Join(templateDir, baseTmplFile)
	cont := path.Join(templateDir, tmplFile)
	t := stdtemplate.Must(stdtemplate.New(baseTmplName).ParseFiles(base, cont))
	return tmpl{Template: t}
}

// Render writes the template out to the response writer (or any errors that
// come up), with values as the template value. If there were errors rendering
// the template, they will not be written out to the writer.
func (t tmpl) Render(w http.ResponseWriter, values interface{}) error {
	var buf bytes.Buffer
	err := t.Template.Execute(&buf, values)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, &buf)
	return err
}
