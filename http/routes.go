package http

import (
	"app/assets"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// setupRoutes for the server.
func (s *Server) setupRoutes() {
	s.mux.Group(func(r chi.Router) {
		r.Use(middleware.Compress(5))
		r.Use(middleware.Logger)

		// Sets up a static file handler with cache busting middleware.
		r.Group(func(r chi.Router) {
			// r.Use(httph.VersionedAssets)

			assets.Static(r)
		})

		r.Get("/", Index())
		r.Post("/savings-form-partial", SavingsFormPartial())
		r.Post("/contact", SavingsFormPartial())
	})
}
