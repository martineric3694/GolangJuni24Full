package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Dalam aplikasi nyata, verifikasi username dan password di database
	if user.Username != "user" && user.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claim{
		Username: user.Username,
		// IsAdmin:  true,
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

func Welcome(c *gin.Context) {
	username := c.GetString("username")

	if c.GetBool("isAdmin") {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome " + username,
			"isAdmin": "Admin",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome " + username,
		})
	}

}
