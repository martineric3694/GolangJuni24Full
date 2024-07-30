package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key") // Ganti dengan key yang lebih kuat

type Claims struct {
	Username string `json:"username"`
	IsAdmin  string `json:"admin"`
	jwt.StandardClaims
}

func handleLogin(c *gin.Context) {
	var user Pengguna
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Dalam aplikasi nyata, verifikasi username dan password di database
	if !checkPengguna(user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user = getDataPengguna(user)

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func handleWelcome(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	unixTimestamp := int64(claims.ExpiresAt)
	t := time.Unix(unixTimestamp, 0)
	formattedTime := t.Format("2006-01-02 15:04:05")

	isAdmin := ""

	if claims.IsAdmin == "1" {
		isAdmin = "Admin"
	} else {
		isAdmin = "Bukan Admin"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Welcome " + claims.Username,
		"admin":       isAdmin,
		"Expired At ": formattedTime,
	})
}
