package router

import (
	"github.com/gin-gonic/gin"
)

func RouterInitialisation(server *gin.Engine) {
	server.Handle("GET", "/events", getEvents)
	server.Handle("GET", "/events/:id", getEvent)
	server.Handle("POST", "/events", createEvent)
	server.Handle("PUT", "/events/:id", updateEvent)
	server.Handle("DELETE", "/events/:id", deleteEvent)
	server.Handle("POST", "/signup", signUp)
}
