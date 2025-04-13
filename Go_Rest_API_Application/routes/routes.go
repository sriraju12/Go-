package routes

import (
	"example.com/restapp/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/events", middleware.Authenticate, createEvents)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", middleware.Authenticate, updateEvent)
	server.DELETE("/events/:id", middleware.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middleware.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middleware.Authenticate, cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
