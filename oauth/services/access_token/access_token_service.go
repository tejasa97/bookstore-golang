package access_token

import (
	"strings"

	"github.com/tejasa97/bookstore-golang/oauth/domain/access_token"
	"github.com/tejasa97/bookstore-golang/oauth/repository/db"
	"github.com/tejasa97/utils-go/rest_errors"

	"github.com/tejasa97/bookstore-golang/oauth/repository/rest"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr)
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *rest_errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(at_request access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr) {
	if err := at_request.Validate(); err != nil {
		return nil, err
	}
	// TODO: Support both grant types, client_credentials and password
	// For now, we only support `password`

	// Authenticate the user against the Users API
	user, err := s.restUsersRepo.LoginUser(at_request.Email, at_request.Password)
	if err != nil {
		return nil, err
	}

	// Generate new access token
	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token to the Cassandra DB
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}
