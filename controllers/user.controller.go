package controllers

import (
	"depmod/models"

	"github.com/labstack/echo"
)

// func FetchAllUser(c echo.Context) error {
// 	result := models.FetchAllUser()
// 	return c.JSON(http.StatusOK, result)
// }

func RegisterUser(c echo.Context) error {
	result, err := models.RegisterUser(
		c.FormValue("username"),
		c.FormValue("email"),
		c.FormValue("password"))
	if err != nil {
		return c.JSON(result.Status, err)
	}
	return c.JSON(result.Status, result)
}
