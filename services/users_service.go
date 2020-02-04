package services

import (
	"github.com/Rizvanas/todolist_users-api/domain/users"
	"github.com/Rizvanas/todolist_users-api/utils/errors"
)

// CreateUser returns created user
func CreateUser(user users.User) (*users.User, *errors.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID returns user stored in database if it exists
func GetUserByID(userID string) (*users.User, *errors.RestError) {
	user, err := users.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
