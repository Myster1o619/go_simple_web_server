package main

import (
	"example.com/rest_api/db"
	"example.com/rest_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()

	router.GET("/events", routes.GetEvents)
	router.GET("/events/:id", routes.GetEventByID)

	router.POST("/events", routes.CreateEvent)

	router.Run(":8080")
}
