package password

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PasswordEncodeSuite struct {
	suite.Suite
}

func (s *PasswordEncodeSuite) SetupTest() {
}

func TestPasswordEncodeAllTests(t *testing.T) {
	suite.Run(t, new(PasswordEncodeSuite))
}
