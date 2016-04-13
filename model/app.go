package model

type App struct {
	Id           int64  `db:"id"`
	ClientId     string `db:"client_id"`
	ClientSecret string `db:"client_secret"`
}
