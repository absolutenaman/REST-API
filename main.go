package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/db"
	"rest-api/models"
)

func main() {
	server := gin.Default()
	db.InitDB()
	server.Handle("GET", "/events", getEvents)
	server.Handle("POST", "/events", createEvent)
	err := server.Run("localhost:8080")
	if err != nil {
		return
	}

}

func getEvents(context *gin.Context) {
	rows, err := db.DB.Query(`SELECT * FROM events`)
	if err != nil {
		panic(err)
	}
	var arr []models.Events
	for rows.Next() {
		var event models.Events
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			panic(err)
		}
		arr = append(arr, event)
	}
	defer rows.Close()
	context.JSON(http.StatusOK, arr)
}
func createEvent(context *gin.Context) {
	var event models.Events
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
	}
	event.ID = 1
	event.UserId = 1
	event.Sava()
	context.JSON(http.StatusOK, gin.H{"message": "Event created succesfully", "event": event})

}
