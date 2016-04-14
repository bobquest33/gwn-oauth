package service

import (
	. "github.com/helderfarias/gwn-oauth/log"
	"github.com/helderfarias/gwn-oauth/model"
	oauthmodel "github.com/helderfarias/oauthprovider-go/model"
	"github.com/jmoiron/sqlx"
)

type PGClientStorage struct {
	db *sqlx.DB
}

func (this *PGClientStorage) FindById(id string) *oauthmodel.Client {
	var app model.App

	err := this.db.Get(&app, this.db.Rebind("SELECT id, client_id, client_secret FROM apps WHERE client_id = ?"), id)
	if err != nil {
		Logger.Error("%s", err)
		return nil
	}

	return &oauthmodel.Client{
		ID:          app.Id,
		Name:        app.ClientId,
		Secret:      app.ClientSecret,
		RedirectUri: "http://localhost",
	}
}

func (this *PGClientStorage) FindByCredencials(clientId, clientSecret string) *oauthmodel.Client {
	var app model.App

	err := this.db.Get(&app, this.db.Rebind("SELECT id, client_id, client_secret FROM apps WHERE client_id = ? and client_secret = ?"), clientId, clientSecret)
	if err != nil {
		Logger.Error("%s", err)
		return nil
	}

	return &oauthmodel.Client{
		ID:          app.Id,
		Name:        app.ClientId,
		Secret:      app.ClientSecret,
		RedirectUri: "http://localhost",
	}
}

func (this *PGClientStorage) Save(entity *oauthmodel.Client) error {
	return nil
}

func (this *PGClientStorage) Delete(entity *oauthmodel.Client) error {
	return nil
}
