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

func TestUserRegister(t *testing.T) {
	log.Println("hello")
	discovery.Init()
	mysql.Init()
	redis.Init()

	h := &Handler{}
	req := &user_core.UserRegisterReq{
		Base: &base.BaseReq{},
		UserInfo: &user_core.UserInfo{
			Username: "wayne",
			Password: "123456",
			Province: "test_province",
			City:     "test_city",
			District: "test_district",
			Role:     "normal",
		},
	}

	res, err := h.UserRegister(context.Background(), req)
	if err != nil {
		t.Errorf("Failed to register user, err: %v", err)
	}
	fmt.Println(litter.Sdump(res))
}
