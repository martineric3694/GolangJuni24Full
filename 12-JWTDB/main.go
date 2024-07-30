package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.POST("/login", handleLogin)
	route.GET("/welcome", handleWelcome).Use(Auth())

	route.Run() //port default 8080
}
