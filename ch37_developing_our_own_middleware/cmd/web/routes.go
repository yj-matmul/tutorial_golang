package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tutorial/pkg/config"
	"github.com/tutorial/pkg/handlers"
)

// routes leads appropriate url to handler function
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// chi is useful for using middleware
	mux.Use(middleware.Recoverer)

	// use our own middleware
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
