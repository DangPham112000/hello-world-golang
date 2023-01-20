package main

import (
	"net/http"

	"github.com/DangPham112000/hello-world-golang/pkg/config"
	"github.com/DangPham112000/hello-world-golang/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", http.HandlerFunc(handlers.RepoC.Home))
	mux.Get("/about", http.HandlerFunc(handlers.RepoC.About))

	return mux
}
