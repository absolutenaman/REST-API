package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	allEventsByTheId, err := models.GetAllEventsById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	context.JSON(http.StatusOK, allEventsByTheId)
}
func createEvent(context *gin.Context) {
	var event models.Events
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
	}
	event.Sava()
	context.JSON(http.StatusOK, gin.H{"message": "Event created succesfully", "event": event})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	allEventByTheId, err := models.GetAllEventsById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	var updatedEvent models.Events

	updatedEvent.ID = allEventByTheId.ID
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	allEventByTheId, err := models.GetAllEventsById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	var deletedEventId models.Events
	deletedEventId.ID = allEventByTheId.ID
	err = deletedEventId.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
