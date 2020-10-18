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
	UpdateUser(int64, users.User) (*users.User, *errors.RestErr)
	DeleteUser(userID int64) *errors.RestErr
}
type usersService struct {
}

func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {

	// user ID has to be more than 0
	if userID <= 0 {
		return nil, errors.NewBadRequest("invalid user id")
	}

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

func (s *usersService) UpdateUser(userID int64, user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := users.DAO.Update(userID, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) DeleteUser(userID int64) *errors.RestErr {

	if err := users.DAO.Delete(userID); err != nil {
		return err
	}

	return nil
}
