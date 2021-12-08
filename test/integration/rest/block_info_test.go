package rest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/olegrom32/tech-assignment/test"
)

const (
	blockNumberInfoURL = "/blocks/%d"
	blockHashInfoURL   = "/blocks/%s"
	blockLatestInfoURL = "/blocks/latest"
)

type BlockInfoTestSuite struct {
	suite.Suite
	container *test.Container
	tearDown  func()
}

func TestBlockInfoTestSuite(t *testing.T) {
	suite.Run(t, new(BlockInfoTestSuite))
}

func (s *BlockInfoTestSuite) SetupTest() {
	s.container, s.tearDown = test.SetUp()
}

func (s *BlockInfoTestSuite) TearDownTest() {
	s.tearDown()
}

func (s *BlockInfoTestSuite) TestSuccessNumber() {
	_ = s.container.RESTClient.
		Get(fmt.Sprintf(blockNumberInfoURL, 9968630)).
		Expect(s.T()).
		Status(http.StatusOK).
		JSON(`{"number":9968630,"hash":"0x1144bc98f856324ab410f1f767478db5e5905614e31ed6ce5dac0ee52273fe28","gas_used":7355272,"transactions_count":37,"timestamp":"0001-01-01T00:00:00Z"}`).
		Done()
}

func (s *BlockInfoTestSuite) TestSuccessHash() {
	_ = s.container.RESTClient.
		Get(fmt.Sprintf(blockHashInfoURL, "0x1144bc98f856324ab410f1f767478db5e5905614e31ed6ce5dac0ee52273fe28")).
		Expect(s.T()).
		Status(http.StatusOK).
		JSON(`{"number":9968630,"hash":"0x1144bc98f856324ab410f1f767478db5e5905614e31ed6ce5dac0ee52273fe28","gas_used":7355272,"transactions_count":37,"timestamp":"0001-01-01T00:00:00Z"}`).
		Done()
}

func (s *BlockInfoTestSuite) TestSuccessLatest() {
	_ = s.container.RESTClient.
		Get(blockLatestInfoURL).
		Expect(s.T()).
		Status(http.StatusOK).
		JSON(`{"number":9968630,"hash":"0x1144bc98f856324ab410f1f767478db5e5905614e31ed6ce5dac0ee52273fe28","gas_used":7355272,"transactions_count":37,"timestamp":"0001-01-01T00:00:00Z"}`).
		Done()
}
