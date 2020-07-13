package main

import (
	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {


	r := gin.Default()
	r.GET("/", homePage)
	r.Run(":3000")
}