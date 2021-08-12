package main

import (
	"github.com/config"
	m "github.com/middlewares"
	"github.com/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	//implement logger middleware
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
