package logger

import (
	"fmt"
	"time"
)

func logger(options *EntranceEntity, level string, message string) {
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

func (s EntranceEntity) Info(log string) {
	logger(&s, "info", log)
}

func (s EntranceEntity) Error(log string) {
	logger(&s, "error", log)
}

func (s EntranceEntity) Success(log string) {
	logger(&s, "success", log)
}

func (s EntranceEntity) Warning(log string) {
	logger(&s, "warning", log)
}
