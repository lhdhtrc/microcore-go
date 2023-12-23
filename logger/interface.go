package logger

type Abstraction interface {
	Info(log string)
	Error(log string)
	Success(log string)
	Warning(log string)
	Mysql(log string)
	Mongo(log string)
}
