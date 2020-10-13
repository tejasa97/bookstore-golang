package users

import (
	"fmt"

	"github.com/tejasa97/bookstore_users-api/utils/date_utils"
	"github.com/tejasa97/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFound(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.Lastname = result.Lastname
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequest(fmt.Sprintf("email %s already registered", current.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("user %d already exists", user.Id))
	}

	now := date_utils.GetNowString()
	user.DateCreated = now

	usersDB[user.Id] = user
	return nil
}
