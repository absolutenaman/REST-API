package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/db"
	"rest-api/router"
)

func main() {
	db.InitDB()
	server := gin.Default()
	router.RouterInitialisation(server)
	err := server.Run("localhost:8080")
	if err != nil {
		return
	}
}
