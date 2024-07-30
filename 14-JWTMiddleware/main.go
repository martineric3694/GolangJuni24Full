package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.POST("/login", Login)
	authorized := route.Group("/data", ValidateJWT())
	{
		authorized.GET("/welcome", Welcome)
		authorized.GET("/adalah", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK BANGET",
			})
		})
	}

	route.Run()
}
