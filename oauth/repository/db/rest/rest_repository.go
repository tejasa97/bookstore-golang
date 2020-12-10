package rest

import (
	"github.com/tejasa97/bookstore-golang/oauth/domain/users"
	"github.com/tejasa97/bookstore-golang/oauth/utils/errors"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	return nil, nil
}
