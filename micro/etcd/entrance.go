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

	Logger logger.Abstraction
}

func New(options *EntranceEntity) micro.Abstraction {
	options.Ctx, options.Cancel = context.WithCancel(context.Background())

	options.Config.MaxRetry = options.Config.MaxRetry | 5
	options.Config.TTL = options.Config.TTL | 5

	options.RetryCount = 0

	return options
}
