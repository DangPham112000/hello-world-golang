package handlers

import (
	"net/http"

	"github.com/DangPham112000/hello-world-golang/pkg/config"
	"github.com/DangPham112000/hello-world-golang/pkg/models"
	"github.com/DangPham112000/hello-world-golang/pkg/render"
)

// Repo the repository used by the handlers
var RepoC *Repository

// Repository is the repository type
type Repository struct {
	AppC *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(aC *config.AppConfig) *Repository {
	return &Repository{
		AppC: aC,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(rC *Repository) {
	RepoC = rC
}

// Home is a home page handler
func (*Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (*Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
