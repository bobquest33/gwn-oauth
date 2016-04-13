package dao

import (
	"encoding/base64"

	"github.com/helderfarias/gwn-oauth/model"
	"github.com/jmoiron/sqlx"
)

const (
	QRY_GET_CONNECTION_BY_APP = `
		SELECT c.type, 
		       c.config,
		       a.client_id,
		       a.client_secret,
		       a.public_key,
		       a.private_key
		  FROM connections c 
		 INNER JOIN apps a ON a.id = c.app_id 
		 WHERE a.client_id = ?
	`
)

type ConnectionDao interface {
	FindByClientId(clientId string) (model.Connection, error)
}

type connectionDao struct {
	db *sqlx.DB
}

func NewConnectionDao(db *sqlx.DB) ConnectionDao {
	return &connectionDao{db: db}
}

func (this *connectionDao) FindByClientId(clientId string) (model.Connection, error) {
	type ConnectionDb struct {
		Type         int         `db:"type"`
		Config       interface{} `db:"config"`
		ClientId     string      `db:"client_id"`
		ClientSecret string      `db:"client_secret"`
		PublicKey    string      `db:"public_key"`
		PrivateKey   string      `db:"private_key"`
	}

	var conn ConnectionDb

	err := this.db.Get(&conn, this.db.Rebind(QRY_GET_CONNECTION_BY_APP), clientId)
	if err != nil {
		return model.Connection{}, err
	}

	var publicKey []byte
	var privateKey []byte

	if conn.PublicKey != "" {
		publicKey, err = base64.StdEncoding.DecodeString(conn.PublicKey)
		if err != nil {
			return model.Connection{}, err
		}
	}

	if conn.PrivateKey != "" {
		privateKey, err = base64.StdEncoding.DecodeString(conn.PrivateKey)
		if err != nil {
			return model.Connection{}, err
		}
	}

	return model.Connection{
		Type:   conn.Type,
		Config: conn.Config,
		App: model.App{
			ClientId:     conn.ClientId,
			ClientSecret: conn.ClientSecret,
			PublicKey:    publicKey,
			PrivateKey:   privateKey,
		},
	}, nil
}
