package config

type DatabaseConf struct {
	DbType     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     int
	DbName     string
	Showsql    bool
}

type EthereumConf struct {
	Endpoint string
}

type LogConf struct {
	Path  string
	Level string
}

type Config struct {
	Database DatabaseConf
	Ethereum EthereumConf
	Log      LogConf
}
