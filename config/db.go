package config

import (
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
)

// DbOptions - used for
type DbOptions struct {
	DB       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"database"`
}

var engine *xorm.Engine

// InitMySQLXorm ..
func InitMySQLXorm(addr string, showSQL bool) {
	e, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		glog.Fatalf("create mysql connection failed. err = %v", err)
	}
	e.ShowSQL(showSQL)
	engine = e
}

// SetMySQLStore ..
func SetMySQLStore(e *xorm.Engine) {
	engine = e
}

// GetMySQLStore ..
func GetMySQLStore() *xorm.Engine {
	return engine
}
