package users

import (
	"net/http"
	"strings"

	"github.com/Rizvanas/todolist_users-api/domain/users"
	"github.com/Rizvanas/todolist_users-api/services"
	"github.com/Rizvanas/todolist_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser creates user
func CreateUser(c *gin.Context) {

	var userCreationRequest users.User

	if err := c.ShouldBindJSON(&userCreationRequest); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	userResult, saveErr := services.CreateUser(userCreationRequest)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}

	c.JSON(http.StatusCreated, userResult)
}

// GetUserByID returns user
func GetUserByID(c *gin.Context) {
	userID := strings.TrimSpace(c.Param("user_id"))

	if userID == "" {
		err := errors.NewBadRequestError("No user id provided.")
		c.JSON(err.StatusCode, err)
		return
	}

	user, err := services.GetUserByID(userID)

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// SearchUser returns user by given criteria
func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "User search is not yet implemented.")
}
