package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/helderfarias/oauthprovider-go"
	oauthgrant "github.com/helderfarias/oauthprovider-go/grant"
	oauthhttp "github.com/helderfarias/oauthprovider-go/http"
	oauthmodel "github.com/helderfarias/oauthprovider-go/model"
	oauthtoken "github.com/helderfarias/oauthprovider-go/token"
	oauthutil "github.com/helderfarias/oauthprovider-go/util"

	"github.com/jmoiron/sqlx"

	"github.com/helderfarias/gwn-oauth/dao"
	. "github.com/helderfarias/gwn-oauth/log"
	"github.com/helderfarias/gwn-oauth/model"
	"github.com/helderfarias/gwn-oauth/password"
)

const (
	TOKEN_EXP_SECONDS        = (60 * 60 * 24 * 30)
	CONNECTION_TYPE_DATABASE = 1
)

type TokenService interface {
	Create(req *http.Request, res http.ResponseWriter) (string, error)
}

type tokenService struct {
	db            *sqlx.DB
	connectionDao dao.ConnectionDao
	config        interface{}
	user          model.User
}

func NewTokenService(db *sqlx.DB) TokenService {
	return &tokenService{
		db:            db,
		connectionDao: dao.NewConnectionDao(db),
	}
}

func (this *tokenService) Create(req *http.Request, res http.ResponseWriter) (string, error) {
	request := &oauthhttp.OAuthRequest{HttpRequest: req}
	response := &oauthhttp.OAuthResponse{HttpRequest: req, HttpResponse: res}

	clientId := request.GetAuthorizationBasic()[0]
	if clientId == "" {
		clientId = request.GetClientId()
	}

	conn, err := this.connectionDao.FindByClientId(clientId)
	if err != nil {
		return "", err
	}

	if this.config, err = this.createConfig(conn.Type, conn.Config); err != nil {
		return "", err
	}

	if this.user, err = this.createUserAndRoles(request); err != nil {
		return "", err
	}

	server := oauthprovider.New().AuthorizationServer()
	server.ClientStorage = &PGClientStorage{db: this.db}
	server.ScopeStorage = &PGScopeStorage{user: this.user}
	server.TokenConverter = &oauthtoken.TokenConverterJwt{
		ExpiryTimeInSecondsForAccessToken: TOKEN_EXP_SECONDS,
		PrivateKey:                        conn.App.PrivateKey,
		PayloadHandler:                    this.createPayload,
	}

	server.AddGrant(&oauthgrant.PasswordGrant{Callback: this.findByUser})
	server.AddGrant(&oauthgrant.AuthzCodeGrant{})
	server.AddGrant(&oauthgrant.ClientCredencial{})

	return server.HandlerAccessToken(request, response)
}

func (this *tokenService) createPayload() map[string]interface{} {
	payload := map[string]interface{}{}
	payload["login"] = this.user.Login
	payload["name"] = this.user.Name
	payload["roles"] = this.user.Roles
	return payload
}

func (this *tokenService) findByUser(user, pass string) *oauthmodel.User {
	var encoder password.PasswordEncode

	encoder = &password.PasswordEncodeDefault{}

	if cfg, ok := this.config.(model.ConnectionDatabaseConfig); ok {
		if cfg.PasswordEncode == "jboss7_md5_base64" {
			encoder = &password.JBoss7MD5Hash{}
		}
	}

	if encoder.Equals(pass, this.user.Password) {
		return &oauthmodel.User{}
	}

	return nil
}

func (t *tokenService) createConfig(typeId int, data interface{}) (interface{}, error) {
	Logger.Debug("Connection Type: %d, Config: %s", typeId, data)

	if typeId == CONNECTION_TYPE_DATABASE {
		var dbconfig model.ConnectionDatabaseConfig
		if err := json.Unmarshal([]byte(fmt.Sprintf("%s", data)), &dbconfig); err != nil {
			return "", err
		}
		return dbconfig, nil
	}

	return nil, errors.New("Connection type was not defined")
}

func (this *tokenService) createUserAndRoles(request oauthhttp.Request) (model.User, error) {
	if request.GetGrantType() != oauthutil.OAUTH_PASSWORD {
		return model.User{}, nil
	}

	if cfg, ok := this.config.(model.ConnectionDatabaseConfig); ok {
		clientDB := sqlx.MustOpen(cfg.Driver, cfg.DataSource)
		defer clientDB.Close()

		var user model.User
		err := clientDB.Get(&user, this.db.Rebind(cfg.User), request.GetUserName())
		if err != nil {
			return model.User{}, err
		}

		err = clientDB.Select(&user.Roles, this.db.Rebind(cfg.Roles), request.GetUserName())
		if err != nil {
			return model.User{}, err
		}

		return user, nil
	}

	return model.User{}, errors.New("Database config type was not defined")
}
