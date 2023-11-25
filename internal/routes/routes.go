package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	
	e.GET("/", index)
}

func index (c echo.Context) error {
	return c.Render(http.StatusOK, "application.html", nil)
}
