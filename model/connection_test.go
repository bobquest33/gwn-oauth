package model

import (
	"encoding/json"
)

func (t *ModelSuite) TestShouldEncodeConfig() {
	con := Connection{
		Type: "database",
		Config: ConnectionDatabaseConfig{
			Driver:         "pg",
			DataSource:     "user=gwn_test password=gwn_test00 dbname=gwn_test host=localhost port=5432 sslmode=disable",
			User:           "select password from users where login = ?",
			Roles:          "select name from roles where login = ?",
			PasswordEncode: "jboss7_md5_base64",
		},
	}

	bytes, _ := json.Marshal(con.Config)

	t.Equal(string(bytes), "")
}
