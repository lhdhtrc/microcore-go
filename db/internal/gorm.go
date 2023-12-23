package internal

import (
	"fmt"
	loggers "github.com/lhdhtrc/microservice-go/logger"
	"gorm.io/gorm/logger"
)

type GormWriter struct {
	logger.Writer
	Logger loggers.Abstraction
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer, s loggers.Abstraction) *GormWriter {
	return &GormWriter{
		Writer: w,
		Logger: s,
	}
}

// Printf 格式化打印日志
func (w *GormWriter) Printf(message string, data ...interface{}) {
	w.Logger.Mysql(fmt.Sprintf(message+"\n", data...))
}
