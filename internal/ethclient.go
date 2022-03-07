package internal

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetClient - Get client
func GetClient(addr string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(addr)
	if err != nil {
		return nil, err
	}
	return client, err
}
