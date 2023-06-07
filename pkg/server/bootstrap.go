package server

import (
	"employee-hierarchy-api/internal/config"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/internal/pg"
	"employee-hierarchy-api/pkg/route"
	"log"

	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) string {
	// base
	config.Init()

	// pg connect
	err := pg.Connect()
	if err != nil {
		log.Fatal(err)
	}
	// error code init
	errorcode.Init()

	// route init
	route.Init(e)

	// services
	cfg := config.GetENV()
	return cfg.Port
}
