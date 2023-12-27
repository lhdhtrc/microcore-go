package internal

import (
	"fmt"
	"github.com/lhdhtrc/microservice-go/logger"
	"time"
)

func Logger(options *logger.EntranceEntity, level string, message string) {
	if options.Config.Enable {
		if options.Config.Console {
			fmt.Printf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), level, message)
		}

		if options.Config.UseRemote {
			if options.Remote != nil {
				options.Remote(level, message)
			} else {
				item := []string{level, message}
				options.Temp = append(options.Temp, item)
			}
		}
	}
}
