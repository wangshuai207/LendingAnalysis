package model

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Ctoken struct {
	ID          int       `xorm:"'id' pk autoincr"`
	Name        string    `xorm:"name" json:"name"`
	Symbol      string    `xorm:"symbol" json:"symbol"`
	Address     string    `xorm:"address" json:"address"`
	Decimals    int       `xorm:"decimals" json:"decimals"`
	Comptroller string    `xorm:"comptroller" json:"comptroller"`
	Admin       string    `xorm:"admin" json:"admin"`
	Underlying  string    `xorm:"underlying" json:"underlying"`
	IsCtoken    int       `xorm:"is_ctoken" json:"is_ctoken"`
	Created     time.Time `xorm:"created" json:"created"`
	Updated     time.Time `xorm:"updated" json:"updated"`
}

func (b *Ctoken) TableName() string {
	return "ctoken"
}

func (b *Ctoken) IsExists(address string, db *xorm.Engine) bool {
	has, _ := db.Exist(&Ctoken{Address: address})
	return has
}

func (b *Ctoken) GetCtoken(address string, db *xorm.Engine) Ctoken {
	var ct Ctoken
	db.Where(" address = ? ", address).Get(&ct)
	return ct
}

func (b *Ctoken) Insert(ct Ctoken, db *xorm.Engine) error {
	_, err := db.Insert(&ct)
	return err
}
