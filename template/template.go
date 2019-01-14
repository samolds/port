// Copyright (C) 2018 - 2019 Sam Olds

package template

import (
	"bytes"
	stdtemplate "html/template"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
)

const (
	// make sure to update the definition in template/base.html if this changes
	baseTmplName = "base"
)

var (
	relTemplateDir = "pages/"

	baseTmplFile  = baseTmplName + ".html"
	homeTmplFile  = "home.html"
	linkTmplFile  = "link.html"
	nowTmplFile   = "now.html"
	errorTmplFile = "error.html"

	// the exported templates that are available to render
	Home  = mustParse(homeTmplFile)
	Link  = mustParse(linkTmplFile)
	Now   = mustParse(nowTmplFile)
	Error = mustParse(errorTmplFile)
)

type tmpl struct {
	*stdtemplate.Template
}

func mustParse(tmplFile string) tmpl {
	_, currFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("can't find templates")
	}
	templateDir := filepath.Join(filepath.Dir(currFile), relTemplateDir)

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
