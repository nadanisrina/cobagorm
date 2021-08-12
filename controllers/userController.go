package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/database"
)

//get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, e := database.GetUser(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	result := c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
	return result
}

//get user
func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	result := c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
	return result
}

//create
func CreateUserController(c echo.Context) error {
	//select all user
	createUser, errCreate := database.CreateUsers(c)
	users, errGetUser := database.GetUsers()

	if errCreate != nil || errGetUser != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errCreate.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created successfully",
		"created": createUser,
		"users":   users,
	})

}

//delete
func DeleteUserController(c echo.Context) error {
	//delete user by id
	id, _ := strconv.Atoi(c.Param("id"))
	deleteUser, errDeleteUser := database.DeleteUser(id)

	if errDeleteUser != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errDeleteUser.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user successfully deleted",
		"user":    deleteUser,
	})

}

//create
func UpdateUserController(c echo.Context) error {
	//update
	id, _ := strconv.Atoi(c.Param("id"))
	_, errUpdateUser := database.UpdateUser(c, id)
	getUsers, _ := database.GetUsers()

	if errUpdateUser != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errUpdateUser.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user successfully",
		"user":    getUsers,
	})
}
