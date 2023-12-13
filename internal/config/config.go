package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCahce      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
}
