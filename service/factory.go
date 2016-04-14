package service

import (
	"github.com/jmoiron/sqlx"
)

type ServiceFactory interface {
	GetTokenService() TokenService
	GetAuthzService() AuthzService
}

type serviceFactory struct {
	tokenService TokenService
	authzService AuthzService
}

func NewServiceFactory(db *sqlx.DB) ServiceFactory {
	factory := &serviceFactory{}
	factory.tokenService = NewTokenService(db)
	factory.authzService = NewAuthzService(db)
	return factory
}

func (s *serviceFactory) GetTokenService() TokenService {
	return s.tokenService
}

func (s *serviceFactory) GetAuthzService() AuthzService {
	return s.authzService
}
