package main

import (
	"strings"

	"github.com/arthurc/bokfor/internal"
	"github.com/arthurc/bokfor/internal/config"
	"github.com/arthurc/bokfor/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = &internal.Templates
	routes.AddRoutes(e)
	
	if err := e.Start(strings.Join([]string{config.Host, config.Port}, ":")); err != nil {
		panic(err)
	}
}
