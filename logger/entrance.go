package logger

type Abstraction interface {
	Info(log string)
	Error(log string)
	Success(log string)
	Warning(log string)
	Mysql(log string)
	Mongo(log string)
}

type EntranceEntity struct {
	Temp [][]string // Temp storage logs

	Console bool // Console whether to output logs on the console
	Enable  bool // Enable whether to enable logs

	Remote    func(level string, message string) // Remote storage logs func
	UseRemote bool                               // UseRemote whether to remote logs

}

func Use(config *EntranceEntity) Abstraction {
	return config
}
