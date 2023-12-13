package main

import (
	"fmt"
	"log"

	"github.com/TartuDen/WordsGame/internal/config"
	"github.com/TartuDen/WordsGame/internal/handlers"
)

const port = ":8081"

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

	handlers.NewHandlers(repo)

}
