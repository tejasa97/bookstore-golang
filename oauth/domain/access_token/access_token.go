package access_token

import (
	"strings"
	"time"

	"github.com/tejasa97/bookstore-golang/oauth/utils/errors"
)

const (
	ExpirationTime = 24 * time.Hour
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	if strings.TrimSpace(at.AccessToken) == "" {
		return errors.NewBadRequest("invalid access token id")
	}
	if at.UserID < 0 {
		return errors.NewBadRequest("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequest("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequest("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(ExpirationTime).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
