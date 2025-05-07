package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		errString := fmt.Sprintf("Unable to retrieve events: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEventByID(context *gin.Context) {
	stringEvtID := context.Param("id")
	eventID, err := strconv.ParseInt(stringEvtID, 10, 64)

	if err != nil {
		errString := fmt.Sprintf("Error converting ID %v to integer: %v", stringEvtID, err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	event, err := models.GetEvent(eventID)

	if err != nil {
		errString := fmt.Sprintf("Unable to retrieve event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusOK, *event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		errString := fmt.Sprintf("Unable to create event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	// dummy data
	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		errString := fmt.Sprintf("Unable to create event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	stringEvtID := context.Param("id")
	eventID, err := strconv.ParseInt(stringEvtID, 10, 64)

	if err != nil {
		errString := fmt.Sprintf("Error converting ID %v to integer: %v", stringEvtID, err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	_, err = models.GetEvent(eventID)

	if err != nil {
		errString := fmt.Sprintf("Unable to retrieve event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		errString := fmt.Sprintf("Unable to update event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	updatedEvent.ID = eventID

	err = updatedEvent.Update()

	if err != nil {
		errString := fmt.Sprintf("Unable to update event: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}
