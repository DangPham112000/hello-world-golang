package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	fmt.Println("render template")
	// create a template cache
	pageCache, err := createTeamplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := pageCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTeamplateCache() (map[string]*template.Template, error) {
	pageCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	fullPathPages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return pageCache, err
	}

	// range through all fullPathPages
	for _, fullPathPage := range fullPathPages {
		pageName := filepath.Base(fullPathPage)
		templateSet, err := template.New(pageName).ParseFiles(fullPathPage)
		if err != nil {
			return pageCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return pageCache, err
		}
		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return pageCache, err
			}
		}

		pageCache[pageName] = templateSet
	}

	return pageCache, nil
}
