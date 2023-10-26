package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafaakilll/ent_todo/auth"
)

// Auth function for checking user is authenticated or not.
//
// If not, return 401 status code and "Unauthorized" message.
//
// If user is authenticated, set user_id to context and continue.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
