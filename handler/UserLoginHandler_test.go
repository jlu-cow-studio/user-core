package handler

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/jlu-cow-studio/common/dal/mysql"
	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/common/discovery"
	"github.com/sanity-io/litter"
)

func TestUserLogin(t *testing.T) {
	log.Println("hello")
	discovery.Init()
	mysql.Init()
	redis.Init()

	s := &Handler{}
	req := &user_core.UserLoginReq{
		Username: "wangmei",
		Password: "123456",
	}

	fmt.Println("hello")
	res, _ := s.UserLogin(context.Background(), req)
	fmt.Println(litter.Sdump(res))
}
