package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jlu-cow-studio/common/dal/redis"
	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
	redis_model "github.com/jlu-cow-studio/common/model/dao_struct/redis"
	"github.com/jlu-cow-studio/user-core/biz"
	"github.com/sanity-io/litter"
)

func (h *Handler) UserAuth(ctx context.Context, req *user_core.UserAuthReq) (res *user_core.UserAuthRes, err error) {
	res = &user_core.UserAuthRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	token := req.Base.Token
	info, err := redis.DB.Get(redis.GetUserTokenKey(token)).Result()

	if err != nil {
		log.Printf("[UserAuth] failed to get token from redis: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "401"
		return res, nil
	}

	userInfoCache := &redis_model.UserInfo{}
	if err := json.Unmarshal([]byte(info), userInfoCache); err != nil {
		log.Printf("[UserAuth] failed to unmarshal user info from redis: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "402"
		return res, nil
	}
	fmt.Println("get user info from cache: ", litter.Sdump(userInfoCache))

	userInfo, err := biz.GetUserInfo(userInfoCache.Username)
	if err != nil {
		log.Printf("[UserAuth] failed to get user info from mysql: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "402"
		return res, nil
	}
	fmt.Println("get user info from mysql: ", litter.Sdump(userInfo))

	if err := biz.CheckUserAuthIsNot(userInfo.Uid, mysql_model.RoleNormal); err != nil {
		log.Printf("[UserAuth] failed to check user auth: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "403"
		return res, nil
	}

	if err := biz.UpdateUserRole(userInfo.Uid, req.GetRole()); err != nil {
		log.Printf("[UserAuth] failed to update user role: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "404"
		return res, nil
	}
	userInfo.Role = req.GetRole()

	infos, err := json.Marshal(userInfo)
	if err != nil {
		log.Printf("[UserAuth] failed to marshal user info: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "405"
		return res, nil
	}

	log.Printf("[UserAuth] set redis %v\n", string(infos))
	if err = redis.DB.Set(redis.GetUserTokenKey(token), string(infos), time.Hour*24).Err(); err != nil {
		log.Printf("[UserAuth] failed to set user info to redis: %v\n", err)
		res.Base.Message = err.Error()
		res.Base.Code = "406"
		return res, nil
	}

	res.Base.Code = "200"
	return res, nil
}
