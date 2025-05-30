package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	usrID := context.GetInt64("usrID")

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

	err = event.Register(usrID)

	if err != nil {
		errString := fmt.Sprintf("Unable to register event: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event registered successfully",
		"event":   event,
	})
}

func cancelRegistration(context *gin.Context) {

}
