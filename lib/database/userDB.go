package database

import (
	"fmt"

	"github.com/config"
	"github.com/labstack/echo/v4"
	"github.com/models"
)

func GetUser(id int) (interface{}, error) {
	var user models.Users
	if e := config.DB.Where("id = ?", id).First(&user).Error; e != nil {
		return nil, e
	}
	fmt.Print("getuserdb", user, nil)
	return user, nil
}

func GetUsers() (interface{}, error) {
	var users []models.Users
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUsers(c echo.Context) (interface{}, error) {
	user := models.Users{}
	c.Bind(&user)
	// pass pointer of data to Create
	if e := config.DB.Save(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

//delete
func DeleteUser(id int) (interface{}, error) {
	var user models.Users
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

//update
func UpdateUser(c echo.Context, id int) (interface{}, error) {
	user := &models.Users{}
	name := c.FormValue("name")
	email := c.FormValue("email")
	pass := c.FormValue("password")
	config.DB.First(&user, id)
	user.Name = name
	user.Email = email
	user.Password = pass

	if e := config.DB.Save(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
