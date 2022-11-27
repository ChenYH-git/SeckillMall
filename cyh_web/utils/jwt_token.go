package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	FrontUserExpireDuration = time.Hour
	FrontUserSecKey         = []byte("frontUserToken")
	AdminUserExpireDuration = time.Hour * 2
	AdminUserSecKey         = []byte("adminUserToken")
)

type UserToken struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	Username string `json:"user_name"`
}

func GenToken(Username string, expireDuration time.Duration, secKey []byte) (string, error) {
	user := UserToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			Issuer:    "micro_gin_vue",
		},
		Username: Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(secKey)
}

func AuthToken(tokenString string, secKey []byte) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (key interface{}, err error) {
		return secKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserToken)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
