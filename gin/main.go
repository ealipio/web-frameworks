package main

import "github.com/gin-gonic/gin"

func main() {
	// instance
	g := gin.Default()

	// routes
	//r.GET("/ping", handler)
	registerRoutes(g)

	// run
	g.Run()

}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
