package logger

type Abstraction interface {
	Info(log string)
	Error(log string)
	Success(log string)
	Warning(log string)
}

type ConfigEntity struct {
	Console   bool `json:"console" bson:"console" yaml:"console" mapstructure:"console"`             // Console whether to output logs on the console
	Enable    bool `json:"enable" bson:"enable" yaml:"enable" mapstructure:"enable"`                 // Enable whether to enable logs
	UseRemote bool `json:"use_remote" bson:"use_remote" yaml:"use_remote" mapstructure:"use_remote"` // UseRemote whether to remote logs
}

type EntranceEntity struct {
	Config ConfigEntity                       // Config logger configs
	Temp   [][]string                         // Temp storage logs
	Remote func(level string, message string) // Remote storage logs func
}

func New(config *EntranceEntity) Abstraction {
	return config
}
