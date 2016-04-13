package model

type User struct {
	Name     string   `json:"name" db:"name"`
	Login    string   `json:"login" db:"login"`
	Password string   `json:"password" db:"password"`
	Roles    []string `json:"roles" db:"-"`
}
