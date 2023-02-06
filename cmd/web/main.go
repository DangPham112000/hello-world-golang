package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DangPham112000/hello-world-golang/pkg/config"
	"github.com/DangPham112000/hello-world-golang/pkg/handlers"
	"github.com/DangPham112000/hello-world-golang/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
