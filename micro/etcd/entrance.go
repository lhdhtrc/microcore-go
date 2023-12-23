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

	Ctx   context.Context
	Cli   *clientv3.Client
	Lease clientv3.LeaseID

	Logger logger.Abstraction
}
