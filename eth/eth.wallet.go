package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func ConnectRPC() {
	client = ConnectRPCEndpoint("https://bsc-dataseed1.bnbchain.org")
}

func GetBalanceETHByAddr(address string) (*big.Int, error) {

	addr := common.HexToAddress(address)

	balance, _ := CheckBalance(client, addr)

	return balance, nil
}

func GetBalanceERC20ByAddr(addr string, contractAddr string) (*big.Int, error) {

	address := common.HexToAddress(addr)

	contractAddress := common.HexToAddress(contractAddr)

	balance, _ := CheckERC20Balance(client, contractAddress, address)

	return balance, nil
}
