package main

import (
	"context"
	"fmt"
	"github.com/sidhuiwnl/RestAPI.git/app"
	"os"
	"os/signal"
)

func main() {
	apps := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) // our own context which responds to our own signal like ctrl + c for graceful shutdown

	err := apps.Start(ctx)

	defer cancel()

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
