package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/utils"
)

func signUp(ctx *gin.Context) {
	var u models.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}
	err = u.AddUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Bad request"})
	}
	ctx.JSON(http.StatusCreated, u)
}
func login(ctx *gin.Context) {
	var u models.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}
	err = u.ValidateUser()
	token, err := utils.TokenGeneration(u.Email, u.Id)
	if err != nil {
		ctx.JSON(401, gin.H{"message": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged In Successfully"})
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
