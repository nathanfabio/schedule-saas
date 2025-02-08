package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nathanfabio/schedule-saas/internal/handlers"
)

// AuthRoutes is a group of auth routes
func AuthRoutes(r chi.Router) {
	r.Post("/singup", handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)
}