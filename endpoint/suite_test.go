package endpoint

import "github.com/gin-gonic/gin"
import "github.com/stretchr/testify/suite"
import "testing"

func init() {
	gin.SetMode(gin.TestMode)
}

type EndpointSuite struct {
	suite.Suite
}

func (s *EndpointSuite) SetupTest() {
}

func TestEndpointAllTests(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}
