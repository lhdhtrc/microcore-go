package internal

import (
	"fmt"
	"github.com/lhdhtrc/microservice-go/logger"
	"time"
)

func Logger(config *logger.EntranceEntity, level string, message string) {
	if config.Enable {
		if config.Console {
			fmt.Printf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), level, message)
		}

		if config.Remote != nil {
			config.Remote(level, message)
		}
	}
}
