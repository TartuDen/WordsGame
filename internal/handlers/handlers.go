package handlers

import (
	"net/http"

	"github.com/TartuDen/WordsGame/internal/config"
	"github.com/TartuDen/WordsGame/internal/models"
	"github.com/TartuDen/WordsGame/internal/renderer"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func GetMainHandler(w http.ResponseWriter, r *http.Request) {
	renderer.RendererTemplate(w, "main.page.html", &models.TemplateData{})
}
