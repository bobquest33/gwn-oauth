package model

import (
	"encoding/json"
)

func (t *ModelSuite) TestShouldEncodeConfig() {
	con := Connection{
		Type: 1,
		Config: ConnectionDatabaseConfig{
			Driver:         "postgres",
			DataSource:     "user=sigem password=sigem00 dbname=sigem host=localhost port=5432 sslmode=disable",
			User:           "select u.senha as password from usuarios u where u.ativo=true and u.login = ?",
			Roles:          "select distinct p.nome from usuarios_grupos ug join usuarios u on ug.id_usuario=u.id join grupos g on ug.id_grupo=g.id join permissoes_grupos pg on pg.id_grupo=g.id join permissoes p on pg.id_permissao=p.id where u.login = ?",
			PasswordEncode: "jboss7_md5_base64",
		},
	}

	bytes, _ := json.Marshal(con.Config)

	t.Equal(string(bytes), "")
}
