package grpc

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type CoreEntity struct {
	Server *grpc.Server
	logger *zap.Logger
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
