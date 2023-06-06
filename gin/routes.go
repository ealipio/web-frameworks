package main

import "github.com/gin-gonic/gin"

func registerRoutes(e *gin.Engine) {

	users := e.Group("/api/v1/users")

	users.GET("/", getAllUsers) // done
	users.POST("/", createUser) // done

	users.GET("/:userId", getUserById)       // done
	users.DELETE("/:userId", deleteUserById) // done
	users.PUT("/:userId", completeUpdateByUserId)
	users.PATCH("/:userId", partialUpdateByUserId)
}
