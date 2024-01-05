package application

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to  start server: %w", err)
	}

	return nil
}
