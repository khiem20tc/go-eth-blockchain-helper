package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func ConnectRPC(rpc string) {
	Client = ConnectRPCEndpoint(rpc)
}

func GetBalanceETHByAddr(address string) (*big.Int, error) {

	addr := common.HexToAddress(address)

	balance, _ := CheckBalance(Client, addr)

	return balance, nil
}

func GetBalanceERC20ByAddr(addr string, contractAddr string) (*big.Int, error) {

	address := common.HexToAddress(addr)

	contractAddress := common.HexToAddress(contractAddr)

	balance, _ := CheckERC20Balance(Client, contractAddress, address)

	return balance, nil
}
