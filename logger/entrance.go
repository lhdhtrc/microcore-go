package logger

type Abstraction interface {
	Info(log string)
	Error(log string)
	Success(log string)
	Warning(log string)
}

type Entity struct {
	Level   uint32
	Message string
}

type ConfigEntity struct {
	Console   bool `json:"console" bson:"console" yaml:"console" mapstructure:"console"`             // Console whether to output logs on the console
	Enable    bool `json:"enable" bson:"enable" yaml:"enable" mapstructure:"enable"`                 // Enable whether to enable logs
	UseRemote bool `json:"use_remote" bson:"use_remote" yaml:"use_remote" mapstructure:"use_remote"` // UseRemote whether to remote logs
}

type EntranceEntity struct {
	Config ConfigEntity     // Config logger configs
	Temp   []Entity         // Temp storage logs
	Remote func(log Entity) // Remote storage logs func
}

func New(config *ConfigEntity) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.Config = *config
	return entity
}
