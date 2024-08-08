package application

import (
	"context"
	"fmt"      // This package implements formatted I/O with functions analogous to C's printf and scanf.
	"net/http" // This package provides HTTP client and server implementations.
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler // This struct is for storing the 
	rdb *redis.Client // This struct is for storing our redis client
}

func New() *App { // This is our constructor
	app := &App {
		router: loadRoutes(),
		rdb: redis.NewClient(&redis.Options{}),
	}
	return app
}

// Creating a reciever (Owner of the method)
func (a *App) Start(ctx context.Context) error{ 
	server := &http.Server {
		Addr: ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err() // Calls the Ping method of the redis client once there is an error (Part of the redis API)
	if err != nil {
		return fmt.Errorf("failed to start connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1) // A channel allows communication across Go routines
	// 1 specifies the buffer size of our channel, only 1 value is coming from this channel
	
	// Go Routine!
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err) // We are publishing the value on the channel
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10) // Without timeout, the shutdown could last indefinitely
		defer cancel()

		return server.Shutdown(timeout) // Graceful Shutdown - Tells the client to finish up any operations before shutting down/terminating the server i.e. handling http requests or writing to a database
	}

	// return nil
}