package users

import (
	"strings"

	"github.com/Rizvanas/todolist_users-api/utils/errors"
)

// User struct
type User struct {
	ID          string `json:"id"`
	FirstName   string `josn:"first_name"`
	LastName    string `josn:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate checks if user object adheres to domain constraints
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address.")
	}

	return nil
}
