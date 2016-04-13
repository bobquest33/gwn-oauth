package model

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ModelSuite struct {
	suite.Suite
}

func (s *ModelSuite) SetupTest() {
}

func TestModelAllTests(t *testing.T) {
	suite.Run(t, new(ModelSuite))
}
