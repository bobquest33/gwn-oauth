package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/helderfarias/oauthprovider-go"
	oauthgrant "github.com/helderfarias/oauthprovider-go/grant"
	oauthhttp "github.com/helderfarias/oauthprovider-go/http"
	oauthmodel "github.com/helderfarias/oauthprovider-go/model"
	"github.com/jmoiron/sqlx"

	. "github.com/helderfarias/gwn-oauth/log"
	"github.com/helderfarias/gwn-oauth/model"
	"github.com/helderfarias/gwn-oauth/password"
)

type TokenService interface {
	Create(req *http.Request) (string, error)
}

type tokenService struct {
	db    *sqlx.DB
	user  model.User
	roles []model.Role
}

func NewTokenService(db *sqlx.DB) TokenService {
	return &tokenService{db: db}
}

func (t *tokenService) Create(req *http.Request) (string, error) {
	request := &oauthhttp.OAuthRequest{HttpRequest: req}

	var conn model.Connection
	err := t.db.Get(&conn, t.db.Rebind("SELECT type, config FROM connections"))
	if err != nil {
		Logger.Error("%s", err)
	}

	var dbConfig model.ConnectionDatabaseConfig
	if err := json.Unmarshal([]byte(fmt.Sprintf("%s", conn.Config)), &dbConfig); err != nil {
		Logger.Error("%s", err)
	} else {
		conn.Config = dbConfig
	}

	clientDB := sqlx.MustOpen(dbConfig.Driver, dbConfig.DataSource)
	defer clientDB.Close()

	var user model.User
	err = clientDB.Get(&user, t.db.Rebind(dbConfig.User), request.GetUserName())
	if err != nil {
		Logger.Error("Error get user: %s", err)
	}

	var roles []model.Role
	err = clientDB.Select(&roles, t.db.Rebind(dbConfig.Roles), request.GetUserName())
	if err != nil {
		Logger.Error("Error get roles: %s", err)
	}

	t.user = user

	server := oauthprovider.New().AuthorizationServer()
	server.ClientStorage = &PGClientStorage{db: t.db}
	server.ScopeStorage = &PGScopeStorage{roles: roles}
	server.AddGrant(&oauthgrant.PasswordGrant{Callback: t.findByUser})

	return server.IssueAccessToken(request)
}

func (t *tokenService) findByUser(user, pass string) *oauthmodel.User {
	encoder := &password.JBoss7MD5Hash{}

	if encoder.Equals(pass, t.user.Password) {
		return &oauthmodel.User{}
	}

	return nil
}
