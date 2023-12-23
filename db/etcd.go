package db

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

func (s EntranceEntity) SetupEtcd(config *ConfigEntity) *clientv3.Client {
	logPrefix := "setup etcd"
	s.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := clientv3.Config{
		DialTimeout: 5 * time.Second,
		Endpoints:   strings.Split(config.Address, ","),
		Context:     ctx,
	}

	switch config.Auth {
	case 1: // not auth
		break
	case 2: // account password
		clientOptions.Username = config.Account
		clientOptions.Password = config.Password
		break
	case 3: // tls
		if config.Tls.CaCert == "" {
			s.Logger.Error(fmt.Sprintf("%s %s", logPrefix, "no CA certificate found"))
			return nil
		}

		if config.Tls.ServerCert == "" {
			s.Logger.Error(fmt.Sprintf("%s %s", logPrefix, "no server certificate found"))
			return nil
		}

		if config.Tls.ServerCertKey == "" {
			s.Logger.Error(fmt.Sprintf("%s %s", logPrefix, "no server certificate key found"))
			return nil
		}

		tlsInfo := transport.TLSInfo{
			CertFile:      config.Tls.ServerCert,
			KeyFile:       config.Tls.ServerCertKey,
			TrustedCAFile: config.Tls.CaCert,
		}

		tlsConfig, err := tlsInfo.ClientConfig()
		if err != nil {
			s.Logger.Error(err.Error())
			return nil
		}

		clientOptions.TLS = tlsConfig
		break
	}

	if config.Mode { // cluster
		clientOptions.AutoSyncInterval = 15 * time.Second
	}

	cli, err := clientv3.New(clientOptions)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
		return nil
	}

	s.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	return cli
}
