package main

import (
	"cashier-system/internal/core/service/cashiersrv"
	"cashier-system/internal/handlers/cashierhdl"
	"cashier-system/internal/repositories/cashierrepo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Customers API
// @version 1.0
// @description.markdown
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	cashierRepository := cashierrepo.NewLocalConfig()
	cashierService := cashiersrv.NewCashierService(cashierRepository)
	cashierHandler := cashierhdl.NewHTTPHandler(cashierService)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.POST("/do-order", cashierHandler.DoOrder)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
