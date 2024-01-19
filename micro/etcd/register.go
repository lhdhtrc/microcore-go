package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lhdhtrc/microservice-go/micro"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"strings"
	"time"
)

// Register etcd service register
func (s EntranceEntity) Register(service string) {
	key := fmt.Sprintf("%s/%s/%d", s.Config.Namespace, service, s.Lease)
	val, _ := json.Marshal(micro.ValueEntity{
		Name:      service,
		Endpoints: s.Config.Address,
	})
	_, err := s.Cli.Put(s.Ctx, key, string(val), clientv3.WithLease(s.Lease))
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	s.logger.Info(fmt.Sprintf("register microservice: %s, %s", key, val))
}

// Deregister etcd service deregister
func (s EntranceEntity) Deregister() {
	if _, err := s.Cli.Revoke(s.Ctx, s.Lease); err != nil {
		s.logger.Error(err.Error())
		return
	}
	s.logger.Info("revoke service lease success")

	key := fmt.Sprintf(s.Config.Namespace)
	res, rErr := s.Cli.KV.Get(s.Ctx, key, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if rErr != nil {
		s.logger.Error(rErr.Error())
		return
	}
	for _, item := range res.Kvs {
		if strings.Contains(string(item.Key), strconv.FormatInt(int64(s.Lease), 10)) {
			if _, err := s.Cli.Delete(s.Ctx, key); err != nil {
				s.logger.Error(err.Error())
				continue
			}
		}
	}
	s.logger.Info("deregister service success")
}

// CreateLease etcd create service instance lease
func (s EntranceEntity) CreateLease() {
	logPrefix := "create lease"
	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	if s.Cli == nil {
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, "etcd client not found"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grant, ge := s.Cli.Grant(ctx, int64(s.Config.TTL))
	if ge != nil {
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, ge.Error()))
		return
	}

	kac, ke := s.Cli.KeepAlive(s.Ctx, grant.ID)
	if ke != nil {
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, ke.Error()))
		return
	}

	go func() {
		//for v := range kac {
		//	s.logger.Info(fmt.Sprintf("microservice lease keepalive success, lease %d, ttl %d", v.ID, v.TTL))
		//}
		for range kac {
		}
		if s.RetryCount < s.Config.MaxRetry {
			time.Sleep(5 * time.Second)

			s.RetryCount++
			s.logger.Info(fmt.Sprintf("retry create lease: %d", s.Config.MaxRetry))
			s.CreateLease()
		}
		s.logger.Info("microservice stop lease alive success")
	}()
	s.logger.Info(fmt.Sprintf("Microservice lease id: %d", grant.ID))
	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	s.Lease = grant.ID
}
