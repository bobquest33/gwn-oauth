package password

type PasswordEncode interface {
	Digest(plain string) string
	Equals(plain, encoded string) bool
}
