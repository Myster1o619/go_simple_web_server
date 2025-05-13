package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	// event routes
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)

	// user routes
	router.POST("/signup", createUser)
}