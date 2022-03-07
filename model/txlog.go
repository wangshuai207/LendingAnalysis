package model

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Txlog struct {
	ID          int       `xorm:"'id' pk autoincr"`
	TxHash      string    `xorm:"txHash" json:"txHash"`
	BlockHash   string    `xorm:"blockHash" json:"blockHash"`
	BlockNumber uint64    `xorm:"blockNumber" json:"blockNumber"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (b *Txlog) TableName() string {
	return "tx_logs"
}

func (b *Txlog) IsExists(address string, db *xorm.Engine) bool {
	has, _ := db.Exist(&Ctoken{Address: address})
	return has
}

func (b *Txlog) Insert(ct Txlog, db *xorm.Engine) error {
	_, err := db.Insert(&ct)
	return err
}
