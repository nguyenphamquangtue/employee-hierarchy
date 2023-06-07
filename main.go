package main

import (
	"employee-hierarchy-api/internal/config"
	"employee-hierarchy-api/pkg/server"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Bootstrap things
	server.Bootstrap(e)

	// Start server
	e.Logger.Fatal(e.Start(config.GetENV().Port))
}
