package dominio

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
)

type Token struct {
	Token string `json:"token"`
}

type payload struct {
	Login string `json:"login"`
	Roles string `json:"roles"`
	Name  string `json:"name"`
}

func (t *Token) GetUsuario() *Usuario {
	parts := strings.Split(t.Token, ".")

	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		log.Println(err)
		return nil
	}

	var pd payload
	err = json.Unmarshal(data, &pd)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &Usuario{
		ID:         pd.Login,
		Nome:       pd.Name,
		Permissoes: strings.Split(pd.Roles, ","),
	}
}
