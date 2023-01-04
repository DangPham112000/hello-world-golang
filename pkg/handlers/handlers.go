package handlers

import (
	"net/http"

	"github.com/DangPham112000/hello-world-golang/pkg/render"
)

// Home is a home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
