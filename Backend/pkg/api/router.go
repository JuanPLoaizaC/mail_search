package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func (a *Api) configureRouter() {
	a.Router = chi.NewRouter()
	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.URLFormat)
	a.Router.Use(render.SetContentType(render.ContentTypeJSON))
	a.configureCORS()
	a.configureRoutes()
}

func (a *Api) configureCORS() {
	a.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		// AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func (a *Api) configureRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the Go and ZinSearch Challenge!"})
	})

	a.Router.Get("/search", a.indexedSearchHandler.SearchTermInEmails)
}
