package chainlink

import (
	"fmt"
	"go-eth-blockchain-helper/eth"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func chainlink() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Connect to BSC Mainnet using archive node URL
	archiveUrl := os.Getenv("ARCHIVE_NODE_URL")

	// Connect to the Ethereum RPC
	eth.ConnectRPC(archiveUrl)

	// Define the contract address
	contractAddr := "0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"
	contractAddress := common.HexToAddress(contractAddr)

	// Define the block numbers to query
	blockNumbers := []string{"20720652", "20720730", "20825470"}

	for _, blockNumStr := range blockNumbers {
		// Convert block number from string to *big.Int
		blockNumber := new(big.Int)
		blockNumber.SetString(blockNumStr, 10) // Base 10 conversion

		// Read function at the specific block
		data, err := eth.ReadFuncAtBlock(eth.Client, eth.ChainlinkABIString, contractAddress, "latestRoundData", []interface{}{}, blockNumber)
		if err != nil {
			fmt.Println("Error reading function:", err)
			return
		}

		// Decode the data
		readableData, err := eth.DecodeData(eth.ChainlinkABIString, "latestRoundData", data)
		if err != nil {
			fmt.Println("Error decoding data:", err)
			return
		}

		// Output the result
		fmt.Printf("Data for block %s: %s\n", blockNumStr, readableData[1])
	}
}
