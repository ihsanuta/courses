package token

import (
	"courses/app/model"
	"courses/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenData struct {
	Data model.User `json:"data"`
	jwt.RegisteredClaims
}

type Token interface {
	GenerateTokenJWT(user model.User) (string, error)
	ValidateToken(signedToken string) (TokenData, error)
}

type token struct {
}

func NewToken() Token {
	return &token{}
}

func (t *token) GenerateTokenJWT(user model.User) (string, error) {
	secretKey := config.JwtConfig["jwt_signature"].(string)
	if secretKey == "" {
		return "", fmt.Errorf("%s", "error get signature key")
	}

	exp := time.Now().AddDate(1, 0, 0)
	issue := time.Now()
	claims := TokenData{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(issue),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// encoded string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("%s", "error generate token")
	}

	return tokenString, nil
}

func (t *token) ValidateToken(signedToken string) (TokenData, error) {
	data := TokenData{}
	secretKey := config.JwtConfig["jwt_signature"].(string)
	if secretKey == "" {
		return data, fmt.Errorf("%s", "error get signature key")
	}

	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenData{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return data, err
	}

	claims, ok := token.Claims.(*TokenData)
	if !ok {
		err = errors.New("couldn't parse claims")
		return data, err
	}
	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return data, err
	}

	return *claims, err
}
