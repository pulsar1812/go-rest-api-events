package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authenticate, cancelRegistration)

	// Or we can use Group
	// authenticated := server.Group("/")
	// authenticated.Use(middlewares.Authenticate)
	// authenticated.POST("/events", createEvent)
	// authenticated.PUT("/events/:id", updateEvent)
	// authenticated.DELETE("/events/:id", deleteEvent)
	// authenticated.POST("/events/:id/register", registerForEvent)
	// authenticated.DELETE("/events/:id/register", cancelRegistration)
}
