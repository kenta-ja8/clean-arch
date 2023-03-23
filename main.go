package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kenta-ja8/clean-arch/internal/driver"
)

func main() {
	fmt.Println("Hello, world!")
	ctx := context.Background()
	driver := driver.InitializeDriver(ctx)
	if err := driver.Serve(ctx, ":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Goodby, world!")
}
