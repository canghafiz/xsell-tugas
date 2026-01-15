package middlewares

import (
	"be/helpers"
	"be/models/repositories"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PwMiddleware(db *gorm.DB, userRepo repositories.UserRepo, jwtKey string) gin.HandlerFunc {
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

		// Extract user id
		userIdFloat, ok := result["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			return
		}

		userId := int(userIdFloat)
		if userId == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			return
		}

		// Check User Id in DB
		userIdValid := userRepo.CheckUserIdValid(db, userId)
		if !userIdValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User Id is",
			})
			return
		}

		// Set user data to context
		c.Set("user_id", userId)
		if userID, exists := result["user_id"]; exists {
			c.Set("user_id", userID)
		}

		c.Next()
	}
}
