package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)



func New() *echo.Echo {
	InitEnvironmentVariables()
	InitDBConnection()
	InitDBCollections()
	echoInstance := echo.New()
	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(middleware.CORS())
	echoInstance.Use(middleware.Recover())
	return echoInstance
}