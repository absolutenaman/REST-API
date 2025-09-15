package router

import (
	"github.com/gin-gonic/gin"
	"rest-api/middlewares"
	"rest-api/models"
	"rest-api/utils"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) AddUser(user *models.User) error {
	return user.AddUser()
}
func (s *UserServiceImpl) ValidateUser(user *models.User) error {
	return user.ValidateUser()
}

type UtilImpl struct {
}

func (u *UtilImpl) TokenGeneration(email string, id int64) (string, error) {
	return utils.TokenGeneration(email, id)
}

type EventsImpl struct {
}

func (e *EventsImpl) GetAllEventsById(id int64) (models.Events, error) {
	return models.GetAllEventsById(id)
}
func (e *EventsImpl) Save(event models.Events) {
	event.Sava()
}

func RouterInitialisation(server *gin.Engine) {
	eventsService := &EventsImpl{}
	userService := &UserServiceImpl{}
	utilService := &UtilImpl{}
	h := NewUserHandler(userService, utilService)
	e := NewEventsHandler(eventsService)
	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", e.createEvent)
	authenticate.PUT("/events/:id", e.updateEvent)
	authenticate.DELETE("/events/:id", e.deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/register", cancellationForEvent)
	server.Handle("GET", "/events", getEvents)
	server.Handle("GET", "/events/:id", e.getEvent)
	server.Handle("POST", "/signup", h.SignUp)
	server.Handle("POST", "/login", h.Login)
}
