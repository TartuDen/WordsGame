package main

import (
	"net/http"

	"github.com/TartuDen/WordsGame/internal/config"
	"github.com/TartuDen/WordsGame/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.GetMainHandler)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
