package services

import (
	"github.com/tejasa97/bookstore-golang/users/domain/users"
	"github.com/tejasa97/utils-go/rest_errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *rest_errors.RestErr)
	CreateUser(users.User) (*users.User, *rest_errors.RestErr)
	UpdateUser(bool, int64, users.User) (*users.User, *rest_errors.RestErr)
	DeleteUser(int64) *rest_errors.RestErr
	FindByStatus(string) (*users.Users, *rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestErr)
}
type usersService struct {
}

func (s *usersService) GetUser(userID int64) (*users.User, *rest_errors.RestErr) {

	// user ID has to be more than 0
	if userID <= 0 {
		return nil, rest_errors.NewBadRequestError("invalid user id")
	}

	user := &users.User{ID: userID}
	err := users.DAO.Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) LoginUser(user_req users.LoginRequest) (*users.User, *rest_errors.RestErr) {
	user := &users.User{
		Email:    user_req.Email,
		Password: user_req.Password,
	}

	if err := users.DAO.GetByEmailPassword(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
	user.Status = users.StatusActive

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.GenerateHashedPassword(); err != nil {
		return nil, err
	}

	if err := users.DAO.Save(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, userID int64, user users.User) (*users.User, *rest_errors.RestErr) {

	if userID <= 0 {
		return nil, rest_errors.NewBadRequestError("invalid user id")
	}

	dbUser := users.User{ID: userID}
	if err := users.DAO.Get(&dbUser); err != nil {
		return nil, err
	}

	// update only provided params if `PATCH` method
	// TODO: Enable update for `password` field also
	if isPartial {
		if user.FirstName != "" {
			dbUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			dbUser.LastName = user.LastName
		}
		if user.Email != "" {
			dbUser.Email = user.Email
		}
		if user.Status != "" {
			dbUser.Status = user.Status
		}
		// update all params if `PUT` method
	} else {
		dbUser.FirstName = user.FirstName
		dbUser.LastName = user.LastName
		dbUser.Email = user.Email
		dbUser.Status = user.Status
	}

	if err := dbUser.Validate(); err != nil {
		return nil, err
	}

	err := users.DAO.Update(&dbUser)
	if err != nil {
		return nil, err
	}

	return &dbUser, nil
}

func (s *usersService) DeleteUser(userID int64) *rest_errors.RestErr {

	if userID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	if err := users.DAO.Delete(userID); err != nil {
		return err
	}

	return nil
}

func (s *usersService) FindByStatus(status string) (*users.Users, *rest_errors.RestErr) {

	var usersFound users.Users
	if err := users.DAO.FindByStatus(&usersFound, status); err != nil {
		return nil, err
	}

	return &usersFound, nil
}
