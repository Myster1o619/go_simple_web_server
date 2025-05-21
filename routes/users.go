package routes

import (
	"fmt"
	"net/http"

	// "strconv"

	"example.com/rest_api/models"
	"example.com/rest_api/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var usr models.User
	err := context.ShouldBindJSON(&usr)

	if err != nil {
		errString := fmt.Sprintf("Unable to parse user data: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	err = usr.Save()

	if err != nil {
		errString := fmt.Sprintf("Unable to create user: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errString,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user":    usr,
	})
}

func login(context *gin.Context) {
	var usr models.User
	err := context.ShouldBindJSON(&usr)

	if err != nil {
		errString := fmt.Sprintf("Unable to parse user data: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errString,
		})
		return
	}

	err = usr.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(usr.Email, usr.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to generate security token",
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": token,
	})
}
