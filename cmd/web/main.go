package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TartuDen/WordsGame/internal/config"
	"github.com/TartuDen/WordsGame/internal/handlers"
	"github.com/TartuDen/WordsGame/internal/renderer"
)

const port = ":8080"

var app config.AppConfig

func main() {
	run()
}

func run() {
	fmt.Println("Starting application of port:", port)

	app.InProduction = false

	repo := handlers.NewRepo(&app)
	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Println(err)
	}
	app.TemplateCache = tc
	app.UseCahce = false
	handlers.NewHandlers(repo)
	renderer.NewTemplate(&app)
	app.InfoLog.Printf("Server is started on port %s",port)

	serv:=&http.Server{
		Addr: port,
		Handler: routes(&app),
	}
	err = serv.ListenAndServe()
	log.Fatal(err)

}
