package routes

import (
	"be-golang/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Auth = middleware.JWT([]byte(os.Getenv("JWT_SECRET")))

var prefix1 = "/api/v1"

func Build(e *echo.Echo) {
	route := e.Group(prefix1)
	// route.Use(helpers.AuthSaas)

	route.POST("/login", controllers.Login)

	route.GET("/barang", controllers.GetListBarang)
	route.GET("/barang/:id", controllers.GetListBarang)
	route.POST("/barang", controllers.CreateBarang)
	route.PUT("/barang/", controllers.UpdateBarang)
	route.DELETE("/barang/:Id", controllers.DeleteBarang)
}
