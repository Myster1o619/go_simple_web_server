package main

import (
	"net/http"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)

	router.Run(":8080")
}
