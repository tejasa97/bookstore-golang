package users

import (
	"strings"
	"time"

	"github.com/tejasa97/utils-go/rest_errors"
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

func (user *User) Validate() *rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Status = strings.TrimSpace(strings.ToLower(user.Status))
	user.Password = strings.TrimSpace(user.Password)

	if user.FirstName == "" {
		return rest_errors.NewBadRequestError("invalid first name")
	}

	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}
	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}

	if (user.Status != StatusActive) && (user.Status != StatusInactive) {
		return rest_errors.NewBadRequestError("invalid status")
	}
	return nil
}

func (user *User) GenerateHashedPassword() *rest_errors.RestErr {

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		hashErr := rest_errors.NewInternalServerError("failed to hash password", err)
		return hashErr
	}

	user.Password = string(hashedPasswordBytes)

	return nil
}

func CompareHashAndPassword(hashedPassword string, password string) error {
	// returns `nil` if equal
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
