package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DangPham112000/hello-world-golang/pkg/config"
	"github.com/DangPham112000/hello-world-golang/pkg/handlers"
	"github.com/DangPham112000/hello-world-golang/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTeamplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repoC := handlers.NewRepo(&app)
	handlers.NewHandlers(repoC)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.RepoC.Home)
	// http.HandleFunc("/about", handlers.RepoC.About)

	fmt.Println(fmt.Sprintf("Starting listening on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
