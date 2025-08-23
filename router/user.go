package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
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
