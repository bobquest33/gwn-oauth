package password

func (t *PasswordEncodeSuite) TestShouldBeEncodeJBoss() {
	encode := JBoss7MD5Hash{}

	t.Equal(encode.Digest("123"), "ICy5YqxZB1uWSwcVLSNLcA==")
}

func (t *PasswordEncodeSuite) TestShouldBeJBossEncodeCompare() {
	encode := JBoss7MD5Hash{}

	t.True(encode.Equals("123", "ICy5YqxZB1uWSwcVLSNLcA=="))
}
