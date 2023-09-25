package main

import (
	"cashier-system/internal/core/service/cashiersrv"
	"cashier-system/internal/handlers/cashierhdl"
	"cashier-system/internal/repositories/cashierrepo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cashierRepository := cashierrepo.NewLocalConfig()
	cashierService := cashiersrv.NewCashierService(cashierRepository)
	cashierHandler := cashierhdl.NewHTTPHandler(cashierService)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	// Routes
	e.POST("/do-order", cashierHandler.DoOrder)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
