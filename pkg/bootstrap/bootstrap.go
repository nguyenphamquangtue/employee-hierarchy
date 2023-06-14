package bootstrap

import (
	"employee-hierarchy-api/internal/config"
	"employee-hierarchy-api/internal/config/errorcode"
	"employee-hierarchy-api/internal/pg"
	"employee-hierarchy-api/pkg/repository"
	"employee-hierarchy-api/pkg/route"
	"employee-hierarchy-api/pkg/service"
	"log"

	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) (pg.DBConnector, string) {
	// base
	config.Init()

	// pg connect
	DBConnector := &pg.PostgreSQLConnector{}
	err := DBConnector.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	// error code init
	errorcode.Init()

	// repository init
	repository := repository.Init(DBConnector)

	// service init
	service := service.Init(repository)

	// route init
	route.Init(e, service)

	// services
	cfg := config.GetENV()
	return DBConnector, cfg.Port
}
