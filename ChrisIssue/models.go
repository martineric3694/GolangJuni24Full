package main

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	Isadmin  int    `json:"isadmin"`
	jwt.StandardClaims
}

type Pengguna struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Isadmin  int    `json:"isadmin"`
}
