package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(email string, userId int64) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
