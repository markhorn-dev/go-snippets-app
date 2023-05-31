package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {

	// Router and middleware setup
	router := chi.NewRouter()
	router.Use(app.recoverPanic)
	router.Use(app.logRequest)
	router.Use(secureHeaders)

	// Serve static files
	fs := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	// And then create the routes using the appropriate methods, patterns and
	// handlers.
	router.Get("/", app.home)
	router.Get("/snippet/view/{id}", app.snippetView)
	router.Get("/snippet/create", app.snippetCreate)
	router.Post("/snippet/create", app.snippetCreatePost)

	// router.HandlerFunc(http.MethodGet, "/", app.home)
	// router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	// router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	// router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	return router
}
