package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate1(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error when parsing template:", err)
		return
	}
}

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in our cache
	_, inMap := templateCache[t]
	if !inMap {
		// need to create the template
		fmt.Println("create cache template")
		err = createTemplateCache(t)
		if err != nil {
			fmt.Println("Error when createTemplateCache:", err)
			return
		}
	} else {
		// we have the template in the cache
		fmt.Println("using cache template")
	}

	tmpl = templateCache[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error when execute template:", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add the template to cache
	templateCache[t] = tmpl

	return nil
}
