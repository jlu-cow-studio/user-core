package main

import (
	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/common/discovery"
	"github.com/jlu-cow-studio/user-core/handler"
	"google.golang.org/grpc"
)

func main() {
	discovery.Init()
	redis.Init()

	s := grpc.NewServer()
	user_core.RegisterUserCoreServiceServer(s, &handler.Handler{})
}
