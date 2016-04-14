package service

import (
	"log"
	"net/http"

	"github.com/helderfarias/gwn-oauth/dao"
	"github.com/helderfarias/gwn-oauth/model"
	"github.com/helderfarias/oauthprovider-go"
	oauthhttp "github.com/helderfarias/oauthprovider-go/http"
	// oauthutil "github.com/helderfarias/oauthprovider-go/util"
	"github.com/jmoiron/sqlx"
)

type AuthzService interface {
	Create(req *http.Request, res http.ResponseWriter) (string, error)
}

type authzService struct {
	db            *sqlx.DB
	connectionDao dao.ConnectionDao
	config        interface{}
	user          model.User
}

func NewAuthzService(db *sqlx.DB) AuthzService {
	return &authzService{
		db:            db,
		connectionDao: dao.NewConnectionDao(db),
	}
}

func (this *authzService) Create(req *http.Request, res http.ResponseWriter) (string, error) {
	request := &oauthhttp.OAuthRequest{HttpRequest: req}
	response := &oauthhttp.OAuthResponse{HttpRequest: req, HttpResponse: res}

	server := oauthprovider.New().AuthorizationServer()
	server.ClientStorage = &PGClientStorage{db: this.db}
	server.ScopeStorage = &PGScopeStorage{user: model.User{}}
	return server.HandlerAuthorize(request, response)
}
