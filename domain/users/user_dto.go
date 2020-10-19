package users

import (
	"strings"
	"time"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tejasa97/bookstore_users-api/utils/errors"
)

type User struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at";"<-:create"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-";sql:"index"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email" gorm:"not_null; unique"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	return nil
}
