package config

import (
	"fmt"

	"github.com/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func InitDB(){
	config := map[string]string{
		"DB_Username" : "root",
		"DB_Password" : "root5638",
		"DB_Port" : "3306",
		"DB_Host" : "127.0.0.1",
		"DB_Name" : "user",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	) 

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!= nil{
		panic(err)
	}

	InitMigrate()
}

func InitMigrate(){
	// AutoMigrate Automatically migrate your schema, to keep your schema up to date.
	DB.AutoMigrate(&models.Users{})
}

