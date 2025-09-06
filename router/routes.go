package router

import (
	"github.com/gin-gonic/gin"
	"rest-api/middlewares"
)

func RouterInitialisation(server *gin.Engine) {
	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/register", cancellationForEvent)

	server.Handle("GET", "/events", getEvents)
	server.Handle("GET", "/events/:id", getEvent)
	server.Handle("POST", "/signup", signUp)
	server.Handle("POST", "/login", login)
}
