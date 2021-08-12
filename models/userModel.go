package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct{
	gorm.Model 
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
