package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectRPCEndpoint(RPCEndpoint string) *ethclient.Client {

	// Connect to Ethereum client
	client, err := ethclient.Dial(RPCEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func VerifyAddress(address string) (bool, error) {
	// Check if the address is a valid Ethereum address
	if !common.IsHexAddress(address) {
		return false, fmt.Errorf("Invalid Ethereum address")
	}

	return true, nil
}

// ERC-20 Token

func CheckERC20Balance(client *ethclient.Client, tokenAddress common.Address, ownerAddress common.Address) (int64, error) {

	balance, err := ReadFunc(client, erc20ABIString, tokenAddress, "balanceOf", []interface{}{ownerAddress})
	balanceStr := hex.EncodeToString(balance)

	// Convert hexadecimal string to decimal
	decimalValue, err := strconv.ParseInt(balanceStr, 16, 64)
	if err != nil {
		fmt.Println("Error converting to decimal:", err)
	}

	return decimalValue, err
}

func CallDataFunction(client *ethclient.Client, toAddress common.Address, value *big.Int) ([]byte, error) {
	// Load the ERC20 ABI (you can replace this with your specific ABI)
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIString)) // Replace with your ERC20 ABI
	if err != nil {
		log.Fatal(err)
	}

	data, err := erc20ABI.Pack("transfer", toAddress, value)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil

}

func EstimateGas(client *ethclient.Client, toAddress common.Address, data []byte) (uint64, error) {
	// Estimate gas
	msg := ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, err
	}

	return gasLimit, nil

}

// Native Token

func CheckBalance(client *ethclient.Client, address common.Address) (*big.Int, error) {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func CreateRawTransaction(client *ethclient.Client, fromAddress common.Address, toAddress common.Address, value *big.Int, data []byte, gasLimit uint64, tip *big.Int, feeCap *big.Int, r *big.Int, s *big.Int, v *big.Int) (*types.Transaction, error) {

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	// ERC-1559
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: feeCap,
		GasTipCap: tip,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
		R:         r,
		S:         s,
		V:         v,
	})

	// tx := types.NewTx(&types.LegacyTx{
	// 	Nonce:    nonce,
	// 	To:       &toAddress,
	// 	Value:    value,
	// 	Gas:      gasLimit,
	// 	GasPrice: feeCap,
	// })

	log.Println("tx", tx)

	return tx, nil
}

func SignRawTransaction(client *ethclient.Client, tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signer := types.NewEIP155Signer(chainID) // Use NewEIP155Signer with the chainID

	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func SendRawTransaction(client *ethclient.Client, signedTx *types.Transaction) error {

	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	fmt.Printf("Transaction hash: %s", signedTx.Hash().Hex())
	return nil
}

func ReadFunc(client *ethclient.Client, ABI string, SCaddress common.Address, funcName string, params []interface{}) ([]byte, error) {
	contract, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		return nil, err
	}

	data, err := contract.Pack(funcName, params...)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &SCaddress,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
