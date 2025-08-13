package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.Handle("GET", "/events", getEvents)

	err := server.Run("localhost:8080")
	if err != nil {
		return
	}

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"events": "eventsList"})
}
