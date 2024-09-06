package main

import (
	"goAssignmentProject/database"
	"goAssignmentProject/router"
)

func main() {
	database.InitDB()

	r := router.SetupRouter()
	r.Run(":8080")
}
