package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginGolang/controller"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/param/:data", GetData)
	r.POST("/param", PostData)
	r.Run(":3000")
}

func GetData(c *gin.Context) {
	data := c.Param("data")
	c.JSON(200, gin.H{
		"data": controller.GetDataAll(data),
	})
}

func PostData(c *gin.Context) {
	var data Data

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":      data.Id + 10,
		"message": data.Message + " sudah terisi",
	})
}

type Data struct {
	Id      uint8  `json:"id"`
	Message string `json:"message"`
}
