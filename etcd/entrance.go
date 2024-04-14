package etcd

import (
	"context"
	"github.com/lhdhtrc/microcore-go/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type prototype struct {
	Config      *model.ConfigEntity
	retryBefore func()
	retryAfter  func()

	retryCount uint32

	ctx    context.Context
	cancel context.CancelFunc

	cli   *clientv3.Client
	lease clientv3.LeaseID

	logger *zap.Logger
}

func New(cli *clientv3.Client, logger *zap.Logger, config *model.ConfigEntity) model.Abstraction {
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
