package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)

	router.Run(":8080")
}