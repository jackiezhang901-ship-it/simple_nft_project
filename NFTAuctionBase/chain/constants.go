package chain

import (
	"strings"

	"github.com/ProjectsTask/EasySwapBase/evm/eip"
	"github.com/pkg/errors"
)

const (
	Eth      = "eth"
	Optimism = "optimism"
	Sepolia  = "sepolia"
)

const (
	EthChainID      = 1        //以太坊主网
	OptimismChainID = 10       //Optimism 主网（Layer 2）
	SepoliaChainID  = 11155111 //以太坊测试网（Sepolia）
)

func UniformAddress(chainName string, address string) (string, error) {
	addr, err := eip.ToCheckSumAddress(address)
	if err != nil {
		return "", errors.Wrap(err, "failed on uniform evm chain address")
	}
	return strings.ToLower(addr), nil

}
