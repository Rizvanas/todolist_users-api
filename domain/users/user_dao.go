package users

import (
	"fmt"

	"github.com/Rizvanas/todolist_users-api/utils/dates"
	"github.com/Rizvanas/todolist_users-api/utils/errors"
)

var (
	usersDB = make(map[string]*User)
)

// Save function saves user to a database
func (user *User) Save() *errors.RestError {
	result := usersDB[user.ID]

	if result != nil {
		return errors.NewBadRequestError("User already exists")
	}

	user.DateCreated = dates.GetUTCTimeString()

	usersDB[user.ID] = user

	return nil
}

// GetUserByID returns user by id
func GetUserByID(userID string) (*User, *errors.RestError) {
	result := usersDB[userID]

	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User %s not found", userID))
	}

	return &User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
	}, nil
}
