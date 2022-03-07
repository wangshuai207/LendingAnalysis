package config

import (
	"fmt"

	"github.com/wangshuai207/LendingAnalysis/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf variable
var Conf *Config

func InitConfig(cfg string) error {

	if err := initViper(cfg); err != nil {
		return err
	}

	return nil
}

func initViper(cfg string) error {
	if cfg != "" {
		viper.SetConfigFile(cfg) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("config") // 设置配置文件路径
		viper.SetConfigName("conf")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量

	if err := getNewConfig(); err != nil {
		return err
	}

	watchConfig()

	return nil
}

func getNewConfig() error {
	var err error
	if err = viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	if Conf, err = NewConfig(viper.GetViper()); err != nil {
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := getNewConfig(); err != nil {
			log.Logger().Error(fmt.Sprintf("getConfigFile Error [%v]", err))
		}
		log.Logger().Info(fmt.Sprintf("Config file changed: %s", e.Name))
	})
}

//Get conf.yaml file to struct
func NewConfig(cfg *viper.Viper) (*Config, error) {
	c := &Config{}

	if err := cfg.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}
