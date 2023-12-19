package handlers

import (
	"net/http"

	"github.com/TartuDen/WordsGame/internal/config"
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

}
