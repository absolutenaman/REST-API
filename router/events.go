package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"strconv"
)

//go:generate mockgen -source=events.go -destination=../mocks/mock_events.go -package=mocks
type EventsService interface {
	GetAllEventsById(event *models.Events) (models.Events, error)
}
type EventsHandler struct {
	eventsService EventsService
}

func NewEventsHandler(service EventsService) *EventsHandler {
	return &EventsHandler{
		eventsService: service,
	}
}
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
	userId := context.GetInt64("userId")
	var event models.Events
	err := context.ShouldBindJSON(&event)
	event.User = userId
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
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
