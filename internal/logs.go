package internal

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetTransactionReceipt - Get Transaction Receipt
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionreceipt
func GetEthLogs(ctx context.Context, client *ethclient.Client, query ethereum.FilterQuery) ([]types.Log, error) {
	res, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
