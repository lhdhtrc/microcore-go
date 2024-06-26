package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lhdhtrc/microcore-go/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Register etcd service register
func (s *MicroEtcdEntity) Register(srv interface{}, desc grpc.ServiceDesc) {
	ref := reflect.TypeOf(srv)
	length := ref.NumMethod()
	for i := 0; i < length; i++ {
		name := ref.Method(i).Name
		key := fmt.Sprintf("%s/%s/%s/%d", s.Config.Namespace, desc.ServiceName, name, s.lease)
		val, _ := json.Marshal(model.ValueEntity{
			Name:      ref.Method(i).Name,
			Endpoints: s.Config.Address,
		})
		_, err := s.cli.Put(s.ctx, key, string(val), clientv3.WithLease(s.lease))
		if err != nil {
			s.logger.Error(err.Error())
			return
		}
		s.logger.Info(fmt.Sprintf("register microservice: %s, %s", key, val))
	}
}

// Deregister etcd service deregister
func (s *MicroEtcdEntity) Deregister() {
	if _, err := s.cli.Revoke(s.ctx, s.lease); err != nil {
		s.logger.Error(err.Error())
		return
	}
	s.logger.Info("revoke service lease success")

	key := fmt.Sprintf(s.Config.Namespace)
	res, rErr := s.cli.KV.Get(s.ctx, key, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if rErr != nil {
		s.logger.Error(rErr.Error())
		return
	}
	for _, item := range res.Kvs {
		if strings.Contains(string(item.Key), strconv.FormatInt(int64(s.lease), 10)) {
			if _, err := s.cli.Delete(s.ctx, key); err != nil {
				s.logger.Error(err.Error())
				continue
			}
		}
	}
	s.logger.Info("deregister service success")
}

// CreateLease etcd create service instance lease
func (s *MicroEtcdEntity) CreateLease() {
	logPrefix := "create lease"
	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	if s.cli == nil {
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, "etcd client not found"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grant, ge := s.cli.Grant(ctx, int64(s.Config.TTL))
	if ge != nil {
		retry(s)
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, ge.Error()))
		return
	}

	kac, ke := s.cli.KeepAlive(s.ctx, grant.ID)
	if ke != nil {
		retry(s)
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, ke.Error()))
		return
	}

	go func() {
		//for v := range kac {
		//	s.logger.Info(fmt.Sprintf("microservice lease keepalive success, lease %d, ttl %d", v.ID, v.TTL))
		//}
		for range kac {
		}
		retry(s)
		s.logger.Info("microservice stop lease alive success")
	}()
	s.logger.Info(fmt.Sprintf("Microservice lease id: %d", grant.ID))
	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	s.lease = grant.ID
}

func (s *MicroEtcdEntity) WithRetryBefore(handle func()) {
	s.retryBefore = handle
}
func (s *MicroEtcdEntity) WithRetryAfter(handle func()) {
	s.retryAfter = handle
}

func retry(s *MicroEtcdEntity) {
	fmt.Println(s.retryCount, s.Config.MaxRetry)
	if s.retryCount < s.Config.MaxRetry {
		if s.retryBefore != nil {
			s.retryBefore()
		}
		time.Sleep(5 * time.Second)

		s.retryCount++
		s.logger.Info(fmt.Sprintf("retry create lease: %d/%d", s.retryCount, s.Config.MaxRetry))
		s.CreateLease()
		if s.retryAfter != nil {
			s.retryAfter()
		}
	}
}
