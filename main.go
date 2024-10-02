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
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Connect to BSC Mainnet using archive node URL
	archiveUrl := os.Getenv("ARCHIVE_NODE_URL")
	client := eth.ConnectRPCEndpoint(archiveUrl)
	fmt.Println("Connected to BSC Mainnet")

	// Define contract address and function parameters
	contractAddress := common.HexToAddress("0x41585C50524fb8c3899B43D7D797d9486AAc94DB")
	blockNumber := big.NewInt(42508758)
	params := []interface{}{
		common.HexToAddress("0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"), // Example token address
		common.HexToAddress("0x1AecbA5Af25d90F9a36Eda909EaA1ED912A891F8"), // Example user address
	}

	// Call the contract function and read data at the specified block
	data, err := eth.ReadFuncAtBlock(
		client,
		eth.ProtocolABIString,
		contractAddress,
		"getUserReserveData",
		params,
		blockNumber,
	)
	if err != nil {
		log.Fatalf("Error calling function: %s", err)
	}

	readableData, err := eth.DecodeData(eth.ProtocolABIString, "getUserReserveData", data)

	if err != nil {
		log.Fatalf("Error decoding data: %s", err)
	}

	fmt.Println(readableData)
}
