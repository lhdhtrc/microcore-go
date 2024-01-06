package download

import "github.com/lhdhtrc/microservice-go/logger"

type EntranceEntity struct {
	Logger logger.Abstraction
}

func New(Logger logger.Abstraction) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.Logger = Logger
	return entity
}
