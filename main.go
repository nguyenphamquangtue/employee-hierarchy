package main

import (
	"context"
	"employee-hierarchy-api/pkg/bootstrap"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Bootstrap things
	dbConnector, port := bootstrap.Bootstrap(e)

	// Start bootstrap
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Failed to start bootstrap: ", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down bootstrap...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Close the database connection
	if err := dbConnector.Close(); err != nil {
		e.Logger.Fatal(err)
	}

	// Shut down the bootstrap
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Failed to shut down bootstrap: ", err)
	}

	fmt.Println("Server gracefully stopped.")
}
