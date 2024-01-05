package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"simple-rest-api-golang/controllers"
	"simple-rest-api-golang/middlewares"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	// Apply middlewares first
	router.Use(middleware.Logger)

	// Define routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Post("/register", controllers.Signup)
	router.Post("/login", controllers.Login)

	// Apply Authentication middleware before defining protected routes
	router.Group(func(chi chi.Router) {
		router.Use(middlewares.Authentication)
		router.Get("/posts", controllers.GetPosts)
	})

	return router
}
