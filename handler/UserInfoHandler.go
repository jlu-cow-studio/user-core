package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	redis_model "github.com/jlu-cow-studio/common/model/dao_struct/redis"
)

func (h *Handler) UserInfo(ctx context.Context, req *user_core.UserInfoReq) (res *user_core.UserInfoRes, err error) {

	res = &user_core.UserInfoRes{
		Base: &base.BaseRes{
			Code: "499",
		},
	}

	cmd := redis.DB.Get(redis.GetUserTokenKey(req.Base.Token))
	if cmd.Err() != nil {
		res.Base.Message = cmd.Err().Error()
		res.Base.Code = "401"
		return res, nil
	}

	info := &redis_model.UserInfo{}
	fmt.Println("get user info :", cmd.Val())

	if err := json.Unmarshal([]byte(cmd.Val()), info); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "402"
		return res, nil
	}

	res.UserInfo = &user_core.UserInfo{
		Username: info.Username,
		Role:     info.Role,
		Province: info.Province,
		City:     info.City,
		District: info.District,
	}

	res.Base.Code = "200"
	return
}
