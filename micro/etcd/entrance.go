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

	Logger logger.Abstraction
}

func Use(config *EntranceEntity) micro.Abstraction {
	config.Ctx, config.Cancel = context.WithCancel(context.Background())

	config.Config.MaxRetry = config.Config.MaxRetry | 5
	config.RetryCount = 0

	return config
}
