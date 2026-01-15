package middlewares

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB, userRepo repositories.UserRepo, jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		// Check Auth Exists
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header required",
			})
			return
		}

		// Check Format Bearer
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			return
		}

		// Extract token from header
		tokenString := strings.TrimPrefix(header, "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is required",
			})
			return
		}

		// Decode Token
		result, errDecode := helpers.DecodeJWT(tokenString, jwtKey)
		if errDecode != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization format",
			})
			return
		}

		// Extract email with type safety
		email, ok := result["email"].(string)
		if !ok || email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			return
		}

		// Check Token in DB
		tokenModel := domains.Users{
			Email: email,
			Token: &tokenString,
		}
		tokenValid := userRepo.CheckTokenValid(db, tokenModel)
		if !tokenValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid or revoked",
			})
			return
		}

		// Set user data to context
		c.Set("user_email", email)
		if userID, exists := result["user_id"]; exists {
			c.Set("user_id", userID)
		}

		c.Next()
	}
}
