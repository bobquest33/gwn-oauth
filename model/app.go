package model

type App struct {
	Id           int64  `db:"id"`
	ClientId     string `db:"client_id"`
	ClientSecret string `db:"client_secret"`
	PublicKey    []byte `db:"public_key"`
	PrivateKey   []byte `db:"private_key"`
}
