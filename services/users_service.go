package services

import (
	"github.com/tejasa97/bookstore_users-api/domain/users"
	"github.com/tejasa97/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
}
type usersService struct {
}

func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	err := users.DAO.Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := users.DAO.Save(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
