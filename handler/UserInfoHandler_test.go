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

func TestUserInfo(t *testing.T) {
	log.Println("hello")
	discovery.Init()
	mysql.Init()
	redis.Init()
	h := &Handler{}
	req := &user_core.UserInfoReq{
		Base: &base.BaseReq{
			Token: "4c62b59b-c809-4987-b10c-f35403227d59",
		},
	}

	res, _ := h.UserInfo(context.Background(), req)

	fmt.Println(litter.Sdump(res))
}
