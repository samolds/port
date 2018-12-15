// Copyright (C) 2018 Sam Olds

package template

import (
	"errors"
	"fmt"
	"html/template"
	//"io/ioutil"
	"net/http"
	//"gopkg.in/webhelp.v1/wherr"
	//"gopkg.in/webhelp.v1/whtmpl"
)

const (
	templateDir = "/Users/samolds/projects/go/src/github.com/samolds/port/template/"
	Home        = "home"
	Links       = "links"
	Now         = "now"
)

var (
	baseTmpl  = templateDir + "base.html"
	homeTmpl  = templateDir + Home + ".html"
	linksTmpl = templateDir + Links + ".html"
	nowTmpl   = templateDir + Now + ".html"
	//masterTemplate = template.New("")

	//// necessary to parse because it's used by other templates
	//baseTemplate = mustParseBase(templateBase, templateDir+"base.html")

	//_ = mustParse(baseTemplate, Home, templateDir+"home.html")
	//_ = mustParse(baseTemplate, Links, templateDir+"links.html")
	//_ = mustParse(baseTemplate, Now, templateDir+"now.html")

	//templates = make(map[string]*template.Template)
	Templates = map[string]*template.Template{
		Home:  template.Must(template.New("base").ParseFiles(baseTmpl, homeTmpl)),
		Links: template.Must(template.New("base").ParseFiles(baseTmpl, linksTmpl)),
		Now:   template.Must(template.New("base").ParseFiles(baseTmpl, nowTmpl)),
	}
)

//func mustParseBase(name, tmplFile string) *template.Template {
//	return template.Must(masterTemplate.New("").ParseFiles(tmplFile))
//}
//
//func mustParse(baseTmpl *template.Template, name,
//	tmplFile string) *template.Template {
//	dupe := template.Must(baseTmpl.Clone())
//	return template.Must(dupe.New(name).ParseFiles(tmplFile))
//}

// Render writes the template out to the response writer (or any errors
// that come up), with value as the template value.
func Render(w http.ResponseWriter, r *http.Request, tmplName string,
	values interface{}) error {
	tmpl, exists := Templates[tmplName]
	if !exists {
		return errors.New(fmt.Sprintf("no template %#v registered", tmplName))
	}

	w.Header().Set("Content-Type", "text/html")
	err := tmpl.Execute(w, values)
	if err != nil {
		return err
	}
	return nil
}
