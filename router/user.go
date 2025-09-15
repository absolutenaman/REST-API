package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
)

//go:generate mockgen -source=user.go -destination=../mocks/mock_user.go -package=mocks

type UserService interface {
	AddUser(user *models.User) error
	ValidateUser(user *models.User) error
}
type Util interface {
	TokenGeneration(email string, id int64) (string, error)
}
type UserHandler struct {
	userService UserService
	util        Util
}

func NewUserHandler(service UserService, util Util) *UserHandler {
	return &UserHandler{
		userService: service,
		util:        util,
	}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var u models.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}
	err = h.userService.AddUser(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Bad request"})
	}
	ctx.JSON(http.StatusCreated, u)
}
func (h *UserHandler) Login(ctx *gin.Context) {
	var u models.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}
	err = h.userService.ValidateUser(&u)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	token, err := h.util.TokenGeneration(u.Email, u.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged In Successfully"})
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
