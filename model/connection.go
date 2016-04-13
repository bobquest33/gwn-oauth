package model

type Connection struct {
	Type   int         `json:"type" db:"type"`
	Config interface{} `json:"config" db:"config"`
	App    App         `json:"app" db:"-"`
}

type ConnectionDatabaseConfig struct {
	Driver         string `json:"driver"`
	DataSource     string `json:"data_source"`
	User           string `json:"user"`
	Roles          string `json:"roles"`
	PasswordEncode string `json:"password_encode"`
}
