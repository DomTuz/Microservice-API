package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App {
		router: loadRoutes(),
	}
	return app
}

// Creating a reciever (Owner of the method)
func (a *App) Start(ctx context.Context) error{ 
	server := &http.Server {
		Addr: ":3000",
		Handler: a.router,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	
	return nil
}