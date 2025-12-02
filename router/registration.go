package router

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=registration.go -destination=../mocks/mock_registration.go -package=mocks
type model interface {
	models.Registrations
}

func registerForEvent(context *gin.Context) {
	eventId := context.Param("id")
	EventIdInInt, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	_, err = models.GetAllEventsById(EventIdInInt)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid event Id"})
		return
	}
	userId := context.GetInt64("userId")
	err = models.Register(userId, EventIdInInt)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Not able to insert the row in registrations table"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "User Registered for event successfully"})
}
func cancellationForEvent(context *gin.Context) {
	eventId := context.Param("id")
	EventIdInInt, err := strconv.ParseInt(eventId, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	userId := context.GetInt64("userId")
	err = models.Cancellation(userId, EventIdInInt)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Error while executing the delete query"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "The registered event has been deleted successfully"})
}
