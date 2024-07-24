package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PreventMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, ok := ctx.Get("x-hasura-role")
		fmt.Println(userRole, ok)
		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		// Check if the user's role is among the allowed roles
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				ctx.Next()
				return
			}
		}
		// If the user's role is not among the allowed roles, return Unauthorized
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		fmt.Println(userRole, ok)

	}
}
