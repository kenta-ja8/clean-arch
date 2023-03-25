package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kenta-ja8/clean-arch/pkg/driver"
)

func main() {
	fmt.Println("Hello, world!")
	ctx := context.Background()
	server, err := driver.InitializeDriver(ctx)
	if err != nil {
		log.Fatal("Failed to initialize driver", err)
	}
	if err := server.Serve(ctx, ":8080"); err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Goodby, world!")
}
