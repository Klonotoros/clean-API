package middleware

import (
	"clean-API/internal/dto"
	"clean-API/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(config dto.Config) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")

		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "No token provided"})
			return
		}

		userId, err := utils.VerifyToken(token, config.SigningSecret)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Wrong token"})
			return
		}

		context.Set("userId", userId)
		context.Next()
	}
}
