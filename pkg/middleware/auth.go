package middleware

import (
	"net/http"
	"os"
	pjwt "real-holat/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	ContextUserID = "user_id"
	ContextRole   = "role"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			ctx.Set(ContextRole, "any")
			ctx.Next()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&pjwt.JWTClaims{},
			func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "token expired or invalid",
				"message": "please login again",
			})
			ctx.Abort()
			return
		}

		claims := token.Claims.(*pjwt.JWTClaims)
		ctx.Set(ContextUserID, claims.UserID)
		ctx.Set(ContextRole, claims.Role)

		ctx.Next()
	}
}
