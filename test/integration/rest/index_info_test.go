package rest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/olegrom32/tech-assignment/test"
)

const (
	indexInfoURL = "/indexes/%d"
)

type IndexInfoTestSuite struct {
	suite.Suite
	container *test.Container
	tearDown  func()
}

func TestIndexInfoTestSuite(t *testing.T) {
	suite.Run(t, new(IndexInfoTestSuite))
}

func (s *IndexInfoTestSuite) SetupTest() {
	s.container, s.tearDown = test.SetUp()
}

func (s *IndexInfoTestSuite) TearDownTest() {
	s.tearDown()
}

func (s *IndexInfoTestSuite) TestSuccess() {
	_ = s.container.RESTClient.
		Get(fmt.Sprintf(indexInfoURL, 1)).
		Expect(s.T()).
		Status(http.StatusOK).
		JSON(`{"name":"DeFi Index (1)","eth_price_in_wei":150000000000000000,"usd_price_in_cents":9500,"usd_capitalization":250000000,"percentage_change":-45}`).
		Done()
}
