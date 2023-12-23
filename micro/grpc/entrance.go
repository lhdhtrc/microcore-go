package grpc

import (
	"context"
	"fmt"
	"github.com/lhdhtrc/microservice-go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
)

type EntranceEntity struct {
	Deploy  bool               `json:"deploy"`  // current service deploy mode: local area network or not
	Address string             `json:"address"` // current service deploy address
	Logger  logger.Abstraction `json:"logger"`  // logger interface
}

func (s EntranceEntity) Dial(endpoint []string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var index int
	length := len(endpoint)
	if length == 0 {
		s.Logger.Warning("no service endpoint are available")
		return nil
	} else if length == 1 {
		index = 0
	} else {
		// Load balancing algorithm that handles multiple nodes
	}

	address := endpoint[index]
	if s.Deploy {
		srv := strings.Split(endpoint[index], ":")
		cur := strings.Split(s.Address, ":")

		if srv[0] == cur[0] {
			address = strings.Join([]string{"127.0.0.1", srv[1]}, ":")
		}
	}

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("the service endpoint is unavailable, error: %s", err.Error()))
		return nil
	}

	return conn
}

func New(config *EntranceEntity) *EntranceEntity {
	entity := new(EntranceEntity)
	entity.Logger = config.Logger

	return entity
}
