package token

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Struct for custom claims
type CustomClaims struct {
	ObjectId string   `json:"objectId"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

func loadSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("JWT_SECRET")
}

func GenerateJWT(objectId string, username, email string, roles []string) (string, error) {
	secretKey := loadSecretKey()

	claims := CustomClaims{
		ObjectId: objectId,
		Username: username,
		Email:    email,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // 3 days expiration
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*CustomClaims, error) {
	secretKey := loadSecretKey()

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetIDFromJWT(tokenString string) (string, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return "", err
	}
	return claims.ObjectId, nil
}
