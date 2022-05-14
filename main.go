package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ymp3cli/ymp3web/routes"
)

func main() {
	e := echo.New()
	e.GET("/", routes.ScanPorts)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Logger.Fatal(e.Start(":1234"))
}
