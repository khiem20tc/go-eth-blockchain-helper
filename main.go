// package main

// import (
// 	eth "go-eth-blockchain-helper/eth"
// )

// func main() {
// 	// Connect Sepolia network
// 	eth.ConnectSepolia()
// }

package main

import (
	"fmt"
	"go-eth-blockchain-helper/eth"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func main() {
	// Connect BSC Mainnet

	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Getting and using a value from .env
	archiveUrl := os.Getenv("ARCHIVE_NODE_URL")

	client := eth.ConnectRPCEndpoint(archiveUrl)

	fmt.Println("Connected to BSC Mainnet", client)

	data, err := eth.ReadFuncAtBlock(
		client,
		eth.ProtocolABIString,
		common.HexToAddress("0x41585C50524fb8c3899B43D7D797d9486AAc94DB"),
		"getReserveData",
		[]interface{}{common.HexToAddress("0x1AecbA5Af25d90F9a36Eda909EaA1ED912A891F8")},
		big.NewInt(42508758),
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
