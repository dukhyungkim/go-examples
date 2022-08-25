package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secret = "powerful_secret"

	testUser = "test_user"
)

func main() {
	tokenString, err := createToken(testUser)
	if err != nil {
		log.Fatalln(err)
	}

	token, err := verifyToken(tokenString)
	if err != nil {
		log.Fatalln(err)
	}

	userId, err := extractUserId(token)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("validate token for %s\n", userId)
}

func createToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, fmt.Errorf("invalidate token")
	}

	return token, nil
}

func extractUserId(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("unexpected claims type")
	}

	userId, has := claims["user_id"].(string)
	if !has {
		return "", fmt.Errorf("failed to find user_id")
	}

	return userId, nil
}
