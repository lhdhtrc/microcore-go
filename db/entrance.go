package db

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/model/base"
)

type ConfigEntity struct {
	Tls base.TLSEntity `json:"tls" bson:"tls" yaml:"tls" mapstructure:"tls"`

	Account  string `json:"account" bson:"account" yaml:"account" mapstructure:"account"`
	Password string `json:"password" bson:"password" yaml:"password" mapstructure:"password"`

	Address  string `json:"address" yaml:"address" mapstructure:"address"`
	Database string `json:"database" yaml:"database" mapstructure:"database"`

	Mode bool `json:"mode" yaml:"mode" mapstructure:"mode"` // Mode is true cluster

	MaxOpenConnects        int  `json:"max_open_connects" bson:"max_open_connects" yaml:"max_open_connects" mapstructure:"max_open_connects"`
	MaxIdleConnects        int  `json:"max_idle_connects" bson:"max_idle_connects" yaml:"max_idle_connects" mapstructure:"max_idle_connects"`
	ConnMaxLifeTime        int  `json:"conn_max_life_time" bson:"conn_max_life_time" yaml:"conn_max_life_time" mapstructure:"conn_max_life_time"`
	SkipDefaultTransaction bool `json:"skip_default_transaction" bson:"skip_default_transaction" yaml:"skip_default_transaction" mapstructure:"skip_default_transaction"`
	PrepareStmt            bool `json:"prepare_stmt" bson:"prepare_stmt" yaml:"prepare_stmt" mapstructure:"prepare_stmt"`

	LoggerEnable bool `json:"logger_enable" bson:"logger_enable" yaml:"logger_enable" mapstructure:"logger_enable"`
}

type EntranceEntity struct {
	logger logger.Abstraction
}

func New(Logger logger.Abstraction) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.logger = Logger
	return entity
}
