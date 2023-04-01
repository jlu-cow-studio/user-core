package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/common/model/http_struct/user"
	"github.com/jlu-cow-studio/user-core/biz"
)

func (h *Handler) Follow(ctx context.Context, req *user_core.FollowReq) (res *user_core.FollowRes, err error) {

	res = &user_core.FollowRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	if req.Action == user.FollowAction_Follow {
		//关注
		if ok, err := biz.CheckUserFollowed(req.FollowerId, req.FollowingId); err != nil {
			res.Base.Message = err.Error()
			res.Base.Code = "401"
			return res, nil
		} else if ok {
			res.Base.Message = "Already followed"
			res.Base.Code = "402"
			return res, nil
		}

		if err := biz.AddUserFollow(req.FollowerId, req.FollowingId); err != nil {
			res.Base.Message = err.Error()
			res.Base.Code = "403"
			return res, nil
		}

	} else if req.Action == user.FollowAction_UnFollow {
		//取消关注

		if ok, err := biz.CheckUserFollowed(req.FollowerId, req.FollowingId); err != nil {
			res.Base.Message = err.Error()
			res.Base.Code = "404"
			return res, nil
		} else if !ok {
			res.Base.Message = "Already unfollowed"
			res.Base.Code = "405"
			return res, nil
		}

		if err := biz.DelUserFollow(req.FollowerId, req.FollowingId); err != nil {
			res.Base.Message = err.Error()
			res.Base.Code = "406"
			return res, nil
		}

	} else {
		res.Base.Message = "unknown action!"
		res.Base.Code = "400"
		return
	}

	return res, nil
}
