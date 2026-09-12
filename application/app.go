package application

import (
	"context"
	"fmt"
	"net/http"
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
		Addr:    ":8080",
		Handler: a.router,
	}

	fmt.Println("Server starting at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	}

	return nil
}
