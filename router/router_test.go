package router

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type routerSuite struct {
	suite.Suite
	client *MockHttpClient
	echo   *MockEcho
	router *Router
}

func (s *routerSuite) SetupTest() {
	s.client = &MockHttpClient{}
	s.echo = &MockEcho{}
	s.router = &Router{
		echo:   s.echo,
		client: s.client,
	}
}

func (s *routerSuite) TearDownTest() {
	s.echo.AssertExpectations(s.T())
	s.client.AssertExpectations(s.T())
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(routerSuite))
}

func (s *routerSuite) TestNewRouter() {
	if r := NewRouter(); s.NotNil(r) {
		s.NotNil(r.echo)
	}
}

func (s *routerSuite) TestStart_ShouldCallEchoStart() {
	address := ":8080"

	s.echo.On("Start", address).Return(nil).Once()

	if err := s.router.Start(address); s.NoError(err) {
	}
}
