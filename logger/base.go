package logger

import "github.com/lhdhtrc/microservice-go/logger/internal"

func (s EntranceEntity) Info(log string) {
	internal.Logger(&s, "info", log)
}

func (s EntranceEntity) Error(log string) {
	internal.Logger(&s, "error", log)
}

func (s EntranceEntity) Success(log string) {
	internal.Logger(&s, "success", log)
}

func (s EntranceEntity) Warning(log string) {
	internal.Logger(&s, "warning", log)
}
