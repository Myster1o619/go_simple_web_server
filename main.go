package main

import (
	"example.com/rest_api/db"
	"example.com/rest_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
