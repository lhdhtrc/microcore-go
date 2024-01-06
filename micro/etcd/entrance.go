package etcd

import (
	"context"
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EntranceEntity struct {
	Config *micro.ConfigEntity

	RetryCount uint32
	Service    *map[string][]string

	Ctx    context.Context
	Cancel context.CancelFunc

	Cli   *clientv3.Client
	Lease clientv3.LeaseID

	logger logger.Abstraction
}

func New(cli *clientv3.Client, logger logger.Abstraction, opt *micro.ConfigEntity) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.Config = opt

	entity.Cli = cli
	entity.logger = logger
	entity.Ctx, entity.Cancel = context.WithCancel(context.Background())

	entity.Config.MaxRetry = entity.Config.MaxRetry | 5
	entity.Config.TTL = entity.Config.TTL | 5

	entity.RetryCount = 0

	return entity
}
