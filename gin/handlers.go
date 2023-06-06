package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

func getAllUsers(g *gin.Context) {
	g.JSON(http.StatusOK, users)
}
func createUser(g *gin.Context) {
	var newUser userType

	err := g.Bind(&newUser)

	if err != nil {
		g.String(http.StatusBadRequest, "Bad Request")
	}

	users = append(users, newUser)
	g.JSON(http.StatusCreated, newUser)
}
func getUserById(g *gin.Context) {
	userId := g.Param("userId")

	user, err := findUserById(userId)
	if err != nil {
		g.String(http.StatusBadRequest, "User Not Found")
	} else {
		g.JSON(http.StatusOK, user)
	}
}
func deleteUserById(g *gin.Context) {
	userId := g.Param("userId")

	index := -1

	for idx, user := range users {
		if user.Id == userId {
			index = idx
			break
		}
	}

	if index == -1 {
		g.String(http.StatusBadRequest, "User Not Found")
	} else {
		users = append(users[:index], users[index+1:]...)
		g.String(http.StatusOK, "User Deleted")
	}

}
func completeUpdateByUserId(g *gin.Context) {
	userId := g.Param("userId")
	userFound, err := findUserById(userId)

	if err != nil {
		g.String(http.StatusBadRequest, "User Not Found")
	} else {
		// Bind
		var userFromBody userType
		err = g.Bind(&userFromBody)

		if err != nil {
			g.String(http.StatusBadRequest, "Bad Request")
		} else {
			// update All Fields
			userFound.Name = userFromBody.Name
			userFound.IsActive = userFromBody.IsActive
			g.String(http.StatusOK, "user updated")
		}
	}

}
func partialUpdateByUserId(g *gin.Context) {
	//userId := g.Param("userId")
	g.String(http.StatusOK, "done")
}

// ------------- utility -------------------------
func findUserById(userId string) (*userType, error) {

	for index, user := range users {

		if user.Id == userId {
			return &users[index], nil
		}
	}

	return nil, errors.New("User Not Found")

}
