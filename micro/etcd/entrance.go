package etcd

import (
	"context"
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type prototype struct {
	Config      *micro.ConfigEntity
	retryBefore func()
	retryAfter  func()

	retryCount uint32

	ctx    context.Context
	cancel context.CancelFunc

	cli   *clientv3.Client
	lease clientv3.LeaseID

	logger logger.Abstraction
}

func New(cli *clientv3.Client, logger logger.Abstraction, config *micro.ConfigEntity) micro.Abstraction {
	ctx, cancel := context.WithCancel(context.Background())

	config.MaxRetry = config.MaxRetry | 5
	config.TTL = config.TTL | 5

	return &prototype{
		Config: config,
		ctx:    ctx,
		cancel: cancel,
		cli:    cli,
		logger: logger,
	}
}
