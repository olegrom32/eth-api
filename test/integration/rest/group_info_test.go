package rest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/olegrom32/tech-assignment/test"
)

const (
	groupInfoURL = "/groups/%d"
)

type GroupInfoTestSuite struct {
	suite.Suite
	container *test.Container
	tearDown  func()
}

func TestGroupInfoTestSuite(t *testing.T) {
	suite.Run(t, new(GroupInfoTestSuite))
}

func (s *GroupInfoTestSuite) SetupTest() {
	s.container, s.tearDown = test.SetUp()
}

func (s *GroupInfoTestSuite) TearDownTest() {
	s.tearDown()
}

func (s *GroupInfoTestSuite) TestSuccess() {
	_ = s.container.RESTClient.
		Get(fmt.Sprintf(groupInfoURL, 12)).
		Expect(s.T()).
		Status(http.StatusOK).
		JSON(`{"name":"DeFi Indexes","indexes":[0,1,2,3,4,5,6]}`).
		Done()
}
