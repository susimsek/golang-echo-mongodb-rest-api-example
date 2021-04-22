package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-echo-mongodb-rest-api-example/controller"
)

func GetSwaggerRoutes(e *echo.Echo) {
	e.GET("/api", controller.RedirectIndexPage)
	e.GET("/api/*", echoSwagger.WrapHandler)
}
