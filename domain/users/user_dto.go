package users

import (
	"fmt"
	"strings"
	"time"

	"github.com/tejasa97/bookstore_users-api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

type Users []User

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

type User struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at";"<-:create"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-";sql:"index"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email" gorm:"not_null; unique"`
	Status    string     `json:"status"`
	Password  string     `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Status = strings.TrimSpace(strings.ToLower(user.Status))
	user.Password = strings.TrimSpace(user.Password)

	if user.FirstName == "" {
		return errors.NewBadRequest("invalid first name")
	}

	if user.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	if user.Password == "" {
		return errors.NewBadRequest("invalid password")
	}

	if (user.Status != StatusActive) && (user.Status != StatusInactive) {
		return errors.NewBadRequest("invalid status")
	}
	return nil
}

func (user *User) GenerateHashedPassword() *errors.RestErr {

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error hashing the input %s", user.Password))
		hashErr := errors.NewInternalServerError("failed to hash password")
		return hashErr
	}

	user.Password = string(hashedPasswordBytes)

	return nil
}
