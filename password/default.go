package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncodeDefault struct {
}

func (e *PasswordEncodeDefault) Digest(plain string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	if err != nil {
		panic(err)
	}

	return string(hashed)
}

func (e *PasswordEncodeDefault) Equals(plain, encoded string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encoded), []byte(plain))
	if err != nil {
		return false
	}

	return true
}
