package etcd

import (
	"context"
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EntranceEntity struct {
	*micro.ConfigEntity
	RetryCount uint32

	Ctx    context.Context
	Cancel context.CancelFunc

	Cli   *clientv3.Client
	Lease clientv3.LeaseID

	Logger logger.Abstraction
}

func New(config *EntranceEntity) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.Logger = config.Logger

	entity.Ctx, entity.Cancel = context.WithCancel(context.Background())

	entity.MaxRetry = entity.MaxRetry | 5
	entity.RetryCount = 0

	return entity
}
