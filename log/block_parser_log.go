package log

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// Init zap log.
func initLog() {
	logger = InitLogger("./block_parser.log", "debug")
}

// Get zap log instance.
func Logger() *zap.Logger {
	once.Do(func() {
		initLog()
	})
	return logger
}
