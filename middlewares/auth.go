package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jtodorovic/macrotrackr/utils"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header"})
		return
	}

	token := strings.TrimPrefix(authHeader, bearerPrefix)

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
