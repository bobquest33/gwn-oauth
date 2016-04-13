package password

func (t *PasswordEncodeSuite) TestShouldBeEncodeDefault() {
	def := PasswordEncodeDefault{}

	pass := def.Digest("123")

	t.NotEmpty(pass)
	t.True(def.Equals("123", "$2a$10$xbR6Z633eV53LX8CLv8TMusD5T5QDv7aG7Ukg2I9clApaEOzZNCeK"))
}
