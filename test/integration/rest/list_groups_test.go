package rest

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/olegrom32/tech-assignment/test"
)

const (
	listGroupsURL = "/groups"
)

type ListGroupsTestSuite struct {
	suite.Suite
	container *test.Container
	tearDown  func()
}

func TestListGroupsTestSuite(t *testing.T) {
	suite.Run(t, new(ListGroupsTestSuite))
}

func (s *ListGroupsTestSuite) SetupTest() {
	s.container, s.tearDown = test.SetUp()
}

func (s *ListGroupsTestSuite) TearDownTest() {
	s.tearDown()
}

func (s *ListGroupsTestSuite) TestSuccess() {
	_ = s.container.RESTClient.
		Get(listGroupsURL).
		Expect(s.T()).
		JSON(`{"groups":[12,13]}`).
		Status(http.StatusOK).
		Done()
}
