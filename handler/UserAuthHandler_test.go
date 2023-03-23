package handler

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/jlu-cow-studio/common/dal/mysql"
	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/common/discovery"
	"github.com/sanity-io/litter"
)

func TestUserAuth(t *testing.T) {
	log.Println("hello")
	discovery.Init()
	mysql.Init()
	redis.Init()

	// user, _ := biz.GetUserInfo("wayne")
	// fmt.Println(litter.Sdump(user))

	s := &Handler{}
	token := "2485ee8f-1a3c-45ed-bf85-79a7caf59c38"
	req := &user_core.UserAuthReq{
		Base: &base.BaseReq{
			Token: token,
		},
		Role: "producer",
	}

	res, _ := s.UserAuth(context.Background(), req)
	fmt.Println(litter.Sdump(res))
}
