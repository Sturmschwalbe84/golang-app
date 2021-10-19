package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World 22")
	})
	r.POST("/", func(c *gin.Context) {
		c.String(200, "Hello World 22")
	})
	r.Run(":8080")
}
