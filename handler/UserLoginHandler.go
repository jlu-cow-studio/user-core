package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	cache "github.com/jlu-cow-studio/common/model/dao_struct/redis"
	"github.com/jlu-cow-studio/user-core/biz"
)

func (h *Handler) UserLogin(ctx context.Context, req *user_core.UserLoginReq) (res *user_core.UserLoginRes, err error) {

	res = &user_core.UserLoginRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	if ok, err := biz.CheckUserExsit(req.Username); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "401"
		return res, nil
	} else if !ok {
		res.Base.Message = "User Not Exsit!"
		res.Base.Code = "402"
		return res, nil
	}

	user, err := biz.GetUserInfo(req.Username)

	if err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "403"
		return res, nil
	}

	if req.Password != user.Password {
		res.Base.Message = "Wrong password!"
		res.Base.Code = "404"
		return res, nil
	}

	token := uuid.NewString()
	info := &cache.UserInfo{
		Uid:      user.Uid,
		Username: user.Username,
		Role:     user.Role,
		Province: user.Province,
		City:     user.City,
		District: user.District,
	}

	infos, err := json.Marshal(info)
	if err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "405"
		return res, nil
	}

	if err = redis.DB.Set(redis.GetUserTokenKey(token), string(infos), time.Hour*24).Err(); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "406"
		return res, nil
	}

	res.Token = token
	res.Base.Code = "200"
	return res, nil
}
