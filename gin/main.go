package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// routes
	r.GET("/ping", handler)
	// run
	r.Run()

}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
