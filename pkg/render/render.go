package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/DangPham112000/hello-world-golang/pkg/config"
	"github.com/DangPham112000/hello-world-golang/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	fmt.Println("render template")
	var pages map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		pages = app.TemplateCache
	} else {
		pages, _ = CreateTeamplateCache()
	}

	// get requested template from cache
	t, ok := pages[tmpl]
	if !ok {
		log.Fatal("Could not get the template from template cache")
	}

	buf := new(bytes.Buffer)

	td = addDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTeamplateCache() (map[string]*template.Template, error) {
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
