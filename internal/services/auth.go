package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(context.Context, string) (string, error)
	ValidateToken(context.Context, string) (string, error)
}

type authService struct {
	jwtKey string
}

type AuthServiceDependency struct {
	JWTKey string
}

func NewAuthService(deps AuthServiceDependency) AuthService {
	return &authService{jwtKey: deps.JWTKey}
}

func (as *authService) GenerateToken(c context.Context, userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Subject:   userID,
	})

	tokenString, err := token.SignedString([]byte(as.jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *authService) ValidateToken(c context.Context, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(as.jwtKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Subject, nil
}
