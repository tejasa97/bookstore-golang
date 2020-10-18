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
	Update(int64, *User) *errors.RestErr
	Delete(int64) *errors.RestErr
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
		return errors.NewBadRequest(fmt.Sprintf("email %s already exists", user.Email))
	}

	return nil
}

func (u *usersDao) Update(userID int64, user *User) *errors.RestErr {

	db_user := User{ID: userID}
	if err := u.Get(&db_user); err != nil {
		return err
	}

	db_user.FirstName = user.FirstName
	db_user.Lastname = user.Lastname
	db_user.Email = user.Email

	if err := users_db.Client.Save(&db_user).Error; err != nil {
		return errors.NewBadRequest(fmt.Sprintf("email %s already exists", user.Email))
	}

	return nil
}

func (u *usersDao) Delete(userID int64) *errors.RestErr {

	result := users_db.Client.Delete(&User{}, userID)
	if result.Error != nil {
		return errors.NewBadRequest(fmt.Sprintf("failed to delete user with id %d", userID))
	}

	if result.RowsAffected == 0 {
		return errors.NewBadRequest(fmt.Sprintf("no user found with id %d", userID))
	}

	return nil
}
