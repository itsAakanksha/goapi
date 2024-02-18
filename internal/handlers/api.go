package handlers

import (
	
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"github.com/itsAakanksha/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}