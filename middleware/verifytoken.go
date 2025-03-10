package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"errors"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(context *gin.Context) {
	TOKEN_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

	var secretKey = []byte(TOKEN_SECRET_KEY)

	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	tokenString := tokenParts[1]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// Set user info in context (assuming claims contain user ID and email)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		email, _ := claims["email"].(string)
		userIDFloat, _ := claims["user_id"].(float64)
		userId := int64(userIDFloat)
		fmt.Print(email)
		fmt.Print(userId)
		// context.Set("userID", userId)
		// context.Set("email", email)
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Allow request to proceed
	context.Next()
}
