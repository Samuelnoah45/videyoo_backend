// middleware/auth.go

package middlewares

import (
	"server/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware checks for a valid token and extracts user information.
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract the token from the "Authorization" header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized "})
			return
		}
		// Parse the token
		tokenWithOutBearer := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer "))
		token, err := jwt.Parse(tokenWithOutBearer, func(token *jwt.Token) (interface{}, error) {
			// Use the same secret key you used for signing the token
			return []byte(config.JWT_SECRET_KEY), nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					// Token is either expired or not valid yet
					ctx.AbortWithStatusJSON(401, gin.H{"message": "Token is expired or not valid yet"})
					return
				}
			} else {
				// Other token validation errors
				ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized something wrong"})
				return
			}
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		// Add user information to the context
		ctx.Set("user-id", claims["https://hasura.io/jwt/claims"].(map[string]interface{})["x-hasura-user-id"])
		// ctx.Set("allowed_roles", claims["https://hasura.io/jwt/claims"].(map[string]interface{})["x-hasura-allowed-roles"])
		ctx.Next()

	}
}
