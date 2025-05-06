package main

import (
	"fmt"
	"net/http"

	"example.com/rest_api/db"
	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()

	router.GET("/events", getEvents)
	router.POST("/events", createEvent)

	router.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		errString := fmt.Sprintf("Unable to create event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	// dummy data
	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}
