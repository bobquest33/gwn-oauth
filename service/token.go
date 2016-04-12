package service

import (
	"net/http"

	"github.com/helderfarias/oauthprovider-go"
	oauthhttp "github.com/helderfarias/oauthprovider-go/http"
	"github.com/helderfarias/oauthprovider-go/model"
	"github.com/jmoiron/sqlx"
)

type TokenService interface {
	Create(req *http.Request) (string, error)
}

type tokenService struct {
	db *sqlx.DB
}

func NewTokenService(db *sqlx.DB) TokenService {
	return &tokenService{db: db}
}

func (t *tokenService) Create(req *http.Request) (string, error) {
	authz := oauthprovider.New().AuthorizationServer()

	authz.ClientStorage.Save(&model.Client{Name: "client", Secret: "client00"})

	return authz.IssueAccessToken(&oauthhttp.OAuthRequest{HttpRequest: req})
}
