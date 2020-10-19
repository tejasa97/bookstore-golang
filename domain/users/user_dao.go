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
	Update(bool, int64, *User) (*User, *errors.RestErr)
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

func (u *usersDao) Update(isPartial bool, userID int64, user *User) (*User, *errors.RestErr) {

	db_user := User{ID: userID}
	if err := u.Get(&db_user); err != nil {
		return nil, err
	}

	// update only provided params if `PATCH` method
	if isPartial {
		if user.FirstName != "" {
			db_user.FirstName = user.FirstName
		}
		if user.LastName != "" {
			db_user.LastName = user.LastName
		}
		if user.Email != "" {
			db_user.Email = user.Email
		}
		// update all params if `PUT` method
	} else {
		db_user.FirstName = user.FirstName
		db_user.LastName = user.LastName
		db_user.Email = user.Email
	}

	if err := db_user.Validate(); err != nil {
		return nil, err
	}

	if err := users_db.Client.Save(&db_user).Error; err != nil {
		return nil, errors.NewBadRequest(fmt.Sprintf("email %s already exists", user.Email))
	}

	return &db_user, nil
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
