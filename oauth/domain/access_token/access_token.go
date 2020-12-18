package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/tejasa97/bookstore-golang/oauth/utils/crypto_utils"
	"github.com/tejasa97/utils-go/rest_errors"
)

const (
	ExpirationTime             = 24 * time.Hour
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	Email    string `json:"email"`
	Password string `json:"password"`

	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *rest_errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return rest_errors.NewBadRequestError("invalid grant type")
	}
	//TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *rest_errors.RestErr {
	if strings.TrimSpace(at.AccessToken) == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID < 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(ExpirationTime).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
