package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", handleLogin)
	router.GET("/welcome", handleWelcome)

	router.Run()
}
