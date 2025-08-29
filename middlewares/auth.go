package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to create an event"})
		return
	}
	userId, err := utils.ValidateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to create an event"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
