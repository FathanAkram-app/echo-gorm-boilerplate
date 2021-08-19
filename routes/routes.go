package routes

import (
	"depmod/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.GET("/getUsers", controllers.FetchAllUser)

	e.POST("/register/user", controllers.RegisterUser)

	return e
}
