package db

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
	"time"
)

func (s EntranceEntity) SetupRedis(config *ConfigEntity) *redis.Client {
	logPrefix := "setup mysql"
	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	// 设置最大超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db, _ := strconv.Atoi(config.Database)
	clientOptions := redis.Options{
		Addr:         config.Address,
		DB:           db,
		MaxIdleConns: config.MaxIdleConnects,
		PoolSize:     config.MaxOpenConnects,
	}

	if config.Account != "" && config.Password != "" {
		clientOptions.Username = config.Account
		clientOptions.Password = config.Password
	}

	if config.Tls.CaCert != "" && config.Tls.ClientCert != "" && config.Tls.ClientCertKey != "" {
		certPool := x509.NewCertPool()
		CAFile, CAErr := os.ReadFile(config.Tls.CaCert)
		if CAErr != nil {
			s.logger.Error(fmt.Sprintf("%s read %s error: %s", logPrefix, config.Tls.CaCert, CAErr.Error()))
			return nil
		}
		certPool.AppendCertsFromPEM(CAFile)

		clientCert, clientCertErr := tls.LoadX509KeyPair(config.Tls.ClientCert, config.Tls.ClientCertKey)
		if clientCertErr != nil {
			s.logger.Error(fmt.Sprintf("%s tls.LoadX509KeyPair err: %v", logPrefix, clientCertErr))
			return nil
		}

		tlsConfig := tls.Config{
			Certificates: []tls.Certificate{clientCert},
			RootCAs:      certPool,
		}
		clientOptions.TLSConfig = &tlsConfig
	}

	cli := redis.NewClient(&clientOptions)

	if c := cli.Ping(ctx); c.Err() != nil {
		s.logger.Error(c.Err().Error())
	}

	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	return cli
}
