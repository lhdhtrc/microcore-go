package logger

import (
	"fmt"
	"time"
)

func logger(options *EntranceEntity, level string, val uint32, message string) {
	if options.Config.Enable {
		if options.Config.Console {
			fmt.Println(fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), level, message))
		}

		temp := Entity{
			Level:   val,
			Message: message,
		}
		if options.Config.UseRemote {
			if options.Remote != nil {
				options.Remote(temp)
			} else {
				options.Temp = append(options.Temp, temp)
			}
		}
	}
}

func (s EntranceEntity) Info(log string) {
	logger(&s, "info", 1, log)
}

func (s EntranceEntity) Error(log string) {
	logger(&s, "error", 2, log)
}

func (s EntranceEntity) Success(log string) {
	logger(&s, "success", 3, log)
}

func (s EntranceEntity) Warning(log string) {
	logger(&s, "warning", 4, log)
}
