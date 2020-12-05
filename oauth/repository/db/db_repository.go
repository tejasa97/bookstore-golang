package db

import (
	"github.com/tejasa97/bookstore-golang/oauth/clients/cassandra"
	"github.com/tejasa97/bookstore-golang/oauth/domain/access_token"
	"github.com/tejasa97/bookstore-golang/oauth/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// TODO: Implement `get access token` from CassandraDB
	return nil, errors.NewInternalServerError("DB connection not implemented")
}
