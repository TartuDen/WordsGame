package renderer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

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

func RendererTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCahce {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cach.")
	}
	td = AddDefaultData(td)

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writting template t obrowser")
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplate))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functionsForTemplate).ParseFiles(page)
		if err!=nil{
			return myCache,err
		}
		matches,err:=filepath.Glob(fmt.Sprintf("%s/*.layout.html",pathToTemplate))
		if err!=nil{
			return myCache,err
		}
		if len(matches)>0{
			ts,err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html",pathToTemplate))
			if err!=nil{
				return myCache, err
			}
		}
		myCache[name]=ts
	}
	return myCache, nil
}
