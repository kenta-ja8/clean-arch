package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kenta-ja8/clean-arch/pkg/external/driver"
)

func main() {
	log.Println("start app")
	ctx := context.Background()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	server, err := driver.InitializeDriver(ctx)
	if err != nil {
		log.Fatal("Failed to initialize driver", err)
	}

	go func() {
		if err := server.Start(ctx); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen %v\n", err)
		}
	}()

	<-sigCh
	log.Println("received SIGINT or SIGTERM")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctxTimeout); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	log.Println("Server stopped")
	log.Println("end app")
}
