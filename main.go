package main

import (
	"GoProductLatihan/database"
	"GoProductLatihan/router"
)

func main() {
	database.Init()

	e := router.Init()

	e.Run(":8080")
}
