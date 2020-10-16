package users

import (
	"fmt"

	"github.com/tejasa97/bookstore_users-api/datasources/sqlite/users_db"
	"github.com/tejasa97/bookstore_users-api/utils/errors"
)

var (
	DAO usersDaoInterface = &usersDao{}
)

type usersDaoInterface interface {
	Get(*User) *errors.RestErr
	Save(*User) *errors.RestErr
}
type usersDao struct {
}

func init() {
	users_db.Client.AutoMigrate(&User{})
}

func (u *usersDao) Get(user *User) *errors.RestErr {

	if users_db.Client.First(&user).Error != nil {
		return errors.NewBadRequest("invalid user id")
	}

	return nil
}

func (u *usersDao) Save(user *User) *errors.RestErr {

	if err := users_db.Client.Create(&user).Error; err != nil {
		return errors.NewBadRequest(fmt.Sprintf("email %s already exists"))
	}

	return nil
}
