package renderer

import (
	"html/template"

	"github.com/TartuDen/WordsGame/internal/config"
	"github.com/TartuDen/WordsGame/internal/models"
)

var app *config.AppConfig

var pathToTemplate = "./template"
var functionsForTemplate = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
}

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}


