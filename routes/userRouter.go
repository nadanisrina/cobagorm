package routes

import (
	"github.com/controllers"
	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"

	mid "github.com/middlewares"
)

func New() *echo.Echo {
	//instance echo
	e := echo.New()

	//bakal return controller
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)

	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(mid.BasicAuthDB))

	return e
}
