package etcd

import (
	"encoding/json"
	"fmt"
	"github.com/lhdhtrc/microservice-go/micro"
	"github.com/lhdhtrc/microservice-go/utils/array"
	clientv3 "go.etcd.io/etcd/client/v3"
	"reflect"
	"strings"
)

// Watcher etcd service watcher
func (s prototype) Watcher(config *[]micro.DiscoverEntity, service *map[string][]string) {
	logPrefix := "[service_endpoint_change] service"
	for _, row := range *config {
		prefix := []string{"/microservice"}
		value := reflect.ValueOf(row)
		for i := 0; i < value.NumField(); i++ {
			v := value.Field(i).String()
			if v != "" {
				prefix = append(prefix, v)
			}
		}

		initService(strings.Join(prefix, "/"), &s, service)

		wc := s.cli.Watch(s.Ctx, strings.Join(prefix, "/"), clientv3.WithPrefix(), clientv3.WithPrevKV())
		go func() {
			for v := range wc {
				for _, e := range v.Events {
					var (
						bytes []byte
						key   string
						val   micro.ValueEntity
					)

					if e.PrevKv != nil {
						key = string(e.PrevKv.Key)
						bytes = e.PrevKv.Value
					} else {
						key = string(e.Kv.Key)
						bytes = e.Kv.Value
					}

					if err := json.Unmarshal(bytes, &val); err != nil {
						s.logger.Warning(err.Error())
						continue
					}

					st := strings.Split(key, "/")
					st = st[:len(st)-1]
					key = strings.Join(st, "/")

					switch e.Type {
					// PUT，新增或替换
					case 0:
						temp := append((*service)[key], val.Endpoints)
						(*service)[key] = array.Unique[string](temp, func(index int, item string) string {
							return item
						})
						s.logger.Success(fmt.Sprintf("%s %s put endpoint, key: %s, endpoint: %s", logPrefix, val.Name, key, val.Endpoints))
					// DELETE
					case 1:
						(*service)[key] = array.Filter((*service)[val.Name], func(index int, item string) bool {
							return item != val.Endpoints
						})
						s.logger.Warning(fmt.Sprintf("%s %s delete endpoint, key: %s, endpoint: %s", logPrefix, val.Name, key, val.Endpoints))
					}
				}
			}
		}()
	}
}

// initService etcd service init
func initService(prefix string, options *prototype, service *map[string][]string) {
	logPrefix := "service discover init service"
	options.logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	res, rErr := options.cli.KV.Get(options.Ctx, prefix, clientv3.WithPrefix())
	if rErr != nil {
		options.logger.Error(fmt.Sprintf("%s %s", logPrefix, rErr.Error()))
		return
	}

	for _, item := range res.Kvs {
		key := string(item.Key)

		var val micro.ValueEntity
		if err := json.Unmarshal(item.Value, &val); err != nil {
			options.logger.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
			return
		}

		st := strings.Split(key, "/")
		st = st[:len(st)-1]
		key = strings.Join(st, "/")

		temp := append((*service)[key], val.Endpoints)
		(*service)[key] = array.Unique[string](temp, func(index int, item string) string {
			return item
		})
	}

	options.logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))
}
