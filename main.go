package main

import (
	"github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Hola, mundo!", "id": id})
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola, mundo!"})
	})
	r.GET("/post/:id", PostMethod)
	r.Run()
}
