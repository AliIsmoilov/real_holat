package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, ok := ctx.Get(ContextRole)
		if !ok {
			role = "any"
		}

		fmt.Printf("Role: %s, Path: %s, Method: %s\n", role, ctx.FullPath(), ctx.Request.Method)
		allowed, err := e.Enforce(
			role.(string),
			ctx.FullPath(),
			ctx.Request.Method,
		)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Allowed: %v\n", allowed)

		if !allowed {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
