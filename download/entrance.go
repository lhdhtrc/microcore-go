package download

import "github.com/lhdhtrc/microservice-go/logger"

type EntranceEntity struct {
	logger logger.Abstraction
}

func New(Logger logger.Abstraction) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.logger = Logger
	return entity
}
