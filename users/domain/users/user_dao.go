package users

import (
	"fmt"

	// "github.com/tejasa97/bookstore-golang/users/domain/users"

	"github.com/tejasa97/bookstore-golang/users/datasources/sqlite/users_db"
	"github.com/tejasa97/utils-go/rest_errors"
)

var (
	DAO usersDaoInterface = &usersDao{}
)

type usersDaoInterface interface {
	Get(*User) *rest_errors.RestErr
	GetByEmailPassword(*User) *rest_errors.RestErr
	Save(*User) *rest_errors.RestErr
	Update(*User) *rest_errors.RestErr
	Delete(int64) *rest_errors.RestErr
	FindByStatus(*Users, string) *rest_errors.RestErr
}
type usersDao struct {
}

func init() {
	users_db.Client.AutoMigrate(&User{})
}

func (u *usersDao) Get(user *User) *rest_errors.RestErr {

	if users_db.Client.First(&user).Error != nil {
		return rest_errors.NewBadRequestError("invalid user id")
	}

	return nil
}

func (u *usersDao) GetByEmailPassword(user *User) *rest_errors.RestErr {

	plainPassword := user.Password

	// only get `active` users
	if users_db.Client.Where("email = ? AND status = ?", user.Email, StatusActive).First(&user).Error != nil {
		return rest_errors.NewBadRequestError("invalid user credentials")
	}

	if err := CompareHashAndPassword(user.Password, plainPassword); err != nil {
		return rest_errors.NewBadRequestError("invalid user credentials")
	}

	return nil
}

func (u *usersDao) Save(user *User) *rest_errors.RestErr {

	if err := users_db.Client.Create(&user).Error; err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}

	return nil
}

func (u *usersDao) Update(user *User) *rest_errors.RestErr {

	if err := users_db.Client.Save(&user).Error; err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
	}

	return nil
}

func (u *usersDao) Delete(userID int64) *rest_errors.RestErr {

	result := users_db.Client.Delete(&User{}, userID)
	if result.Error != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("failed to delete user with id %d", userID))
	}

	if result.RowsAffected == 0 {
		return rest_errors.NewBadRequestError(fmt.Sprintf("no user found with id %d", userID))
	}

	return nil
}

func (u *usersDao) FindByStatus(users *Users, status string) *rest_errors.RestErr {

	results := users_db.Client.Find(&users, User{Status: status})
	if results.Error != nil {
		return rest_errors.NewInternalServerError("error while trying to find users by status", results.Error)
	}

	return nil
}
