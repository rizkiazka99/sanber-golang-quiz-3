package middleware

import (
	"github.com/gin-gonic/gin"
)

func ValidateAccessToken(ctx *gin.Context) (id string, error string) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		return "", "access token is required"
	} else {
		tokenString := authHeader[len("Bearer "):]
		userId, err := ValidateJWT(tokenString)

		if err != nil {
			return "", "invalid or expired access token"
		} else {
			return userId, ""
		}
	}
}
