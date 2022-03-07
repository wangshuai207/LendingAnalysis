package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wangshuai207/LendingAnalysis/internal"
	"github.com/wangshuai207/LendingAnalysis/log"
)

func main() {
	var client *ethclient.Client
	// create  rpc client
	client, err := internal.GetClient("https://mainnet.infura.io/v3/d6f7ae662eb94fbca97980db33db35ae")
	if err != nil {
		log.Logger().Error(fmt.Sprintf("failed to get ethclient: %v", client))
		return
	}
	log.Logger().Info("ethclient succeed")
	//Comptroller :0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(7710733),
		// ToBlock:   new(big.Int).SetInt64(1111),
		// Addresses: []common.Address{
		// 	"",
		// },
		Topics: [][]common.Hash{{common.HexToHash("0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d")}},
	}
	logs, err := internal.GetEthLogs(context.Background(), client, query)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("failed to GetEthLogs: %v", err))
		return
	}
	for _, lg := range logs {
		//common.HexToAddress(lg.Topics[0].Hex())
		//fmt.Println(lg.Topics[0].Hex())
		log.Logger().Info(lg.Address.Hex())
		log.Logger().Info(common.Bytes2Hex(lg.Data))
	}
}
