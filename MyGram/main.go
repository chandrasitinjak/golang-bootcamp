package main

import (
	"final_task/database"
	"final_task/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()

	r.Run(":8080")
}
