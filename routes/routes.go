package routes

import (
	"example.com/rest_api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// event routes
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)

	authenticated := router.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", middleware.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// user routes
	router.POST("/signup", createUser)
	router.POST("/login", login)
}