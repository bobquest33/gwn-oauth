package service

import (
	"github.com/jmoiron/sqlx"
)

type ServiceFactory interface {
	GetTokenService() TokenService
}

type serviceFactory struct {
	tokenService TokenService
}

func NewServiceFactory(db *sqlx.DB) ServiceFactory {
	factory := &serviceFactory{}
	factory.tokenService = NewTokenService(db)
	return factory
}

func (s *serviceFactory) GetTokenService() TokenService {
	return s.tokenService
}
