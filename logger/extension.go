package logger

import "github.com/lhdhtrc/microservice-go/logger/internal"

func (s EntranceEntity) Mysql(log string) {
	internal.Logger(&s, "mysql", log)
}

func (s EntranceEntity) Mongo(log string) {
	internal.Logger(&s, "mongo", log)
}
