package routes

import (
	"fmt"
	"net/http"
	// "strconv"

	"example.com/rest_api/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var usr models.User
	err := context.ShouldBindJSON(&usr)

	if err != nil {
		errString := fmt.Sprintf("Unable to create user: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	// dummy data
	// event.ID = 1
	// event.UserID = 1

	err = usr.Save()

	if err != nil {
		errString := fmt.Sprintf("Unable to create user: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user":   usr,
	})
}