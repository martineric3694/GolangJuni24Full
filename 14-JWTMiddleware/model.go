package main

import "github.com/golang-jwt/jwt"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claim struct {
	Username string
	IsAdmin  bool
	jwt.StandardClaims
}
