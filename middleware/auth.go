package middleware

import (
	"net/http"

	"example.com/rest_api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized User",
		})
		return
	}

	usrID, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized User",
		})
		return
	}

	context.Set("usrID", int64(usrID))
	context.Next()
}
