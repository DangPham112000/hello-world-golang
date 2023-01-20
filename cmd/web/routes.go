package main

import (
	"net/http"

	"github.com/DangPham112000/hello-world-golang/pkg/config"
	"github.com/DangPham112000/hello-world-golang/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.RepoC.Home))
	mux.Get("/about", http.HandlerFunc(handlers.RepoC.About))

	return mux
}
