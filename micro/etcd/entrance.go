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

	Ctx    context.Context
	Cancel context.CancelFunc

	Cli   *clientv3.Client
	Lease clientv3.LeaseID

	logger logger.Abstraction
}

func New(cli *clientv3.Client, logger logger.Abstraction, config *micro.ConfigEntity) *EntranceEntity {
	ctx, cancel := context.WithCancel(context.Background())

	config.MaxRetry = config.MaxRetry | 5
	config.TTL = config.TTL | 5

	return &EntranceEntity{
		Config: config,
		Ctx:    ctx,
		Cancel: cancel,
		Cli:    cli,
		logger: logger,
	}
}
