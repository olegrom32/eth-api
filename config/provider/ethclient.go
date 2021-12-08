package provider

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/olegrom32/tech-assignment/config"
	"github.com/olegrom32/tech-assignment/pkg/contract"
)

func NewETHClient(cfg config.Ropsten) (*ethclient.Client, error) {
	return ethclient.Dial(cfg.URL)
}

func NewContract(ethClient *ethclient.Client, cfg config.Ropsten) (*contract.ContractCaller, error) {
	return contract.NewContractCaller(common.HexToAddress(cfg.Contract), ethClient)
}
