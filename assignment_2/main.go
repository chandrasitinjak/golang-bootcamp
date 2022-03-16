package main

import (
	"assignment_2/database"
	"assignment_2/router"
)

func main() {
	database.StartDB()

	router.StartServer().Run(":8080")

}
