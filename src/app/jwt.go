package app

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID int64  `json:"id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (app *App) GenerateJWT(userID int64, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 24 * 365).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(app.config.jwtSecret))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to sign jwt with error: %v", err))
	}
	return signedToken, nil
}

func (app *App) validateToken(tokenString string, expectedRole string) (*Claims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.config.jwtSecret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid token signature")
		}
		return nil, errors.New(fmt.Sprintf("Failed to parse token with error: %v", err))
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("Failed to extract claims from token")
	}

	if claims.Role != expectedRole {
		return nil, errors.New("Invalid token role")
	}

	return claims, nil
}
