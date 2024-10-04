package main

import (
	"fmt"
	"go-eth-blockchain-helper/eth"
	"log"
)

func main() {
	eth.ConnectRPC("https://bsc-dataseed.binance.org")

	// Get ETH balance by address
	ethBalance, err := eth.GetBalanceETHByAddr("0xa180fe01b906a1be37be6c534a3300785b20d947")

	if err != nil {
		log.Fatalf("Error getting ETH balance: %s", err)
	}

	fmt.Println("ETH balance:", ethBalance)

	// Get ERC20 balance by address
	erc20Balance, err := eth.GetBalanceERC20ByAddr("0xa180fe01b906a1be37be6c534a3300785b20d947", "0x55d398326f99059fF775485246999027B3197955")

	if err != nil {
		log.Fatalf("Error getting ERC20 balance: %s", err)
	}

	fmt.Println("ERC20 balance:", erc20Balance)
}
