package grpc

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strings"
)

type ConfigEntity struct {
	Deploy  bool   `json:"deploy"`  // current service deploy mode: local area network or not
	Address string `json:"address"` // current service deploy address
}

type CoreEntity struct {
	Server *grpc.Server
	logger *zap.Logger
}

func (s *CoreEntity) Dial(endpoint []string, opt *ConfigEntity) *grpc.ClientConn {
	var index int
	length := len(endpoint)
	if length == 0 {
		s.logger.Warn("no service endpoint are available")
		return nil
	} else if length == 1 {
		index = 0
	} else {
		// Load balancing algorithm that handles multiple nodes
	}

	address := endpoint[index]
	if opt.Deploy {
		srv := strings.Split(endpoint[index], ":")
		cur := strings.Split(opt.Address, ":")

		if srv[0] == cur[0] {
			address = strings.Join([]string{"127.0.0.1", srv[1]}, ":")
		}
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Error(fmt.Sprintf("the service endpoint is unavailable, error: %s", err.Error()))
		return nil
	}

	return conn
}

func (s *CoreEntity) CreateServer(handle func(server *grpc.Server), address string) {
	logPrefix := "setup grpc server"
	s.logger.Info(fmt.Sprintf("%s %s %s", logPrefix, address, "start ->"))

	listen, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
		return
	}
	server := grpc.NewServer()

	/*-------------------------------------Register Microservice---------------------------------*/
	handle(server)
	/*-------------------------------------Register Microservice---------------------------------*/

	s.logger.Info(fmt.Sprintf("%s %s", logPrefix, "register server done ->"))
	go func() {
		sErr := server.Serve(listen)
		if sErr != nil {
			s.logger.Error(fmt.Sprintf("%s %s", logPrefix, sErr.Error()))
			return
		}
	}()
	s.Server = server
}

func New(logger *zap.Logger) *CoreEntity {
	return &CoreEntity{logger: logger}
}
