package password

import (
	"crypto/md5"
	"encoding/base64"
)

type JBoss7MD5Hash struct {
}

func (e *JBoss7MD5Hash) Digest(plain string) string {
	digest := md5.New()
	digest.Write([]byte(plain))
	hash := digest.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash)
}

func (e *JBoss7MD5Hash) Equals(plain, encoded string) bool {
	return e.Digest(plain) == encoded
}
