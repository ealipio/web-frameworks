package main

import (
	"net/http"

	"errors"

	"github.com/labstack/echo/v4"
)

func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
func createUser(c echo.Context) error {
	var newUser userType

	err := c.Bind(&newUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	users = append(users, newUser)
	return c.JSON(http.StatusCreated, newUser)
}
func getUserById(c echo.Context) error {
	userId := c.Param("userId")

	for _, user := range users {
		if user.Id == userId {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.String(http.StatusBadRequest, "User Not Found")
}
func deleteUserById(c echo.Context) error {
	userId := c.Param("userId")

	index := -1

	for idx, user := range users {
		if user.Id == userId {
			index = idx
			break
		}
	}

	if index == -1 {
		return c.String(http.StatusBadRequest, "User Not Found")
	}

	users = append(users[:index], users[index+1:]...)
	return c.String(http.StatusOK, "User Deleted")
}
func completeUpdateByUserId(c echo.Context) error {
	userId := c.Param("userId")
	userFound, err := findUserById(userId)

	if err != nil {
		return c.String(http.StatusBadRequest, "User Not Found")
	}

	// Bind
	var userFromBody userType
	err = c.Bind(&userFromBody)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// update All Fields
	userFound.Name = userFromBody.Name
	userFound.IsActive = userFromBody.IsActive
	return c.String(http.StatusOK, "user updated")
}
func partialUpdateByUserId(c echo.Context) error {
	//userId := c.Param("userId")
	return c.String(http.StatusOK, "done")
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
