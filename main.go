package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/DomTuz/Microservice-API/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) // This method takes in a context and a signal, returns a context if there is a signal
	defer cancel() // defer keyword is used to stop a function's execution (In this case, )

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
} 
