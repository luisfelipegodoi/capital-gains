package financial

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceSuite struct {
	suite.Suite
	*require.Assertions
	ctrl    *gomock.Controller
	service CapitalGains
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}
