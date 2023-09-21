package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// configureRouter setups chi router mildelware for http communication
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

// configureCORS setups the router to allow cross-source resource sharing
// only with the VUE client and certain methods.
func (a *Api) configureCORS() {
	a.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

// configureRoutes sets up the routes for the app.
func (a *Api) configureRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"message": "Welcome to the Go and ZinSearch Challenge!"})
	})

	a.Router.Get("/search", a.indexedSearchHandler.SearchTermInEmails)
}
