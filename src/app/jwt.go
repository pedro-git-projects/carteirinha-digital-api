package app

import (
	"errors"
	"fmt"
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
