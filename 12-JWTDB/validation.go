package main

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(tokenString string) (claimsReturn Claims, err error) {
	if tokenString == "" {
		err := errors.New("Token string is empty")
		return claimsReturn, err
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		err := errors.New("Invalid token")
		return claimsReturn, err
	}

	claimsReturn = *claims

	return claimsReturn, nil
}
