package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is a home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error when parsing template:", err)
		return
	}
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
