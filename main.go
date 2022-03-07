package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/wangshuai207/LendingAnalysis/config"
	"github.com/wangshuai207/LendingAnalysis/internal"
	"github.com/wangshuai207/LendingAnalysis/log"
	"github.com/wangshuai207/LendingAnalysis/model"
)

var (
	cfg = pflag.StringP("config", "c", "./config/conf.yaml", " config file path")
)

func initMysqlDB(conf *config.DatabaseConf) {
	var dbOpt config.DbOptions

	dbOpt.DB = conf.DbType
	dbOpt.User = conf.DbUser
	dbOpt.Password = conf.DbPassword
	dbOpt.Host = conf.DbHost
	dbOpt.Port = fmt.Sprintf("%d", conf.DbPort)
	dbOpt.Schema = conf.DbName

	var addr = fmt.Sprint(dbOpt.User, ":", dbOpt.Password, "@(", dbOpt.Host,
		":", dbOpt.Port, ")/", dbOpt.Schema, "?charset=utf8mb4&parseTime=True")
	config.InitMySQLXorm(addr, conf.Showsql)
}

func main() {
	pflag.Parse()

	// Set the time zone to UTC+8.
	time.Local = time.FixedZone("CST", 3600*8)

	// init config via viper
	if err := config.InitConfig(*cfg); err != nil {
		fmt.Printf("failed to InitConfig: %v", err)
		panic(err)
	}

	initMysqlDB(&config.Conf.Database)
	log.Logger().Info("init mysql succeed")
	var client *ethclient.Client
	// create  rpc client
	client, err := internal.GetClient("https://mainnet.infura.io/v3/d6f7ae662eb94fbca97980db33db35ae")
	if err != nil {
		log.Logger().Error(fmt.Sprintf("failed to get ethclient: %v", client))
		return
	}
	log.Logger().Info("ethclient succeed")
	//Comptroller :0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b
	// query := ethereum.FilterQuery{
	// 	FromBlock: new(big.Int).SetInt64(7710733),
	// 	// ToBlock:   new(big.Int).SetInt64(1111),
	// 	// Addresses: []common.Address{
	// 	// 	"",
	// 	// },
	// 	Topics: [][]common.Hash{{common.HexToHash("0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d")}},
	// }
	// logs, err := internal.GetEthLogs(context.Background(), client, query)
	// if err != nil {
	// 	log.Logger().Error(fmt.Sprintf("failed to GetEthLogs: %v", err))
	// 	return
	// }
	// ct := model.Ctoken{}
	// for _, lg := range logs {
	// 	if common.Bytes2Hex(lg.Data) == "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000003d9819210a31b4961b30ef54be2aed79b9c9cd3b" {
	// 		log.Logger().Info(lg.Address.Hex())
	// 		ct.Address = lg.Address.Hex()
	// 		ct.IsCtoken = 1

	// 		ct.Insert(ct, config.GetMySQLStore())
	// 	}
	// }

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(14090831),
		//ToBlock:   &big.Int{},
		Addresses: []common.Address{common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
			common.HexToAddress("0xB3319f5D18Bc0D84dD1b4825Dcde5d5f7266d407"),
			common.HexToAddress("0x6C8c6b02E7b2BE14d4fA6022Dfd6d75921D90E4E"),
			common.HexToAddress("0xA4C993E32876795ABf80842Adb0a241bb0EeCD47"),
			common.HexToAddress("0x35a18000230da775cac24873d00ff85bccded550"),
			common.HexToAddress("0x5d3a536E4D6DbD6114cc1Ead35777bAB948E3643")},
		Topics: [][]common.Hash{{common.HexToHash("0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52")}},
	}
	logs, err := internal.GetEthLogs(context.Background(), client, query)
	if err != nil {
		log.Logger().Error(fmt.Sprintf("failed to GetEthLogs: %v", err))
		return
	}
	tx := model.Txlog{}
	for _, lg := range logs {
		//log.Logger().Info(lg.Address.Hex())
		tx.TxHash = lg.TxHash.Hex()
		tx.BlockHash = lg.BlockHash.Hex()
		tx.BlockNumber = lg.BlockNumber
		//tx.Insert(tx, config.GetMySQLStore())
		fmt.Println("\"" + tx.TxHash + "\",")
	}
}
