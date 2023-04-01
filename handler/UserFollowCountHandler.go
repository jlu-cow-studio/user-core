package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/user-core/biz"
)

func (h *Handler) FollowCount(ctx context.Context, req *user_core.FollowCountReq) (res *user_core.FollowCountRes, err error) {

	res = &user_core.FollowCountRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	if followerCount, followingCount, err := biz.FollowCount(req.UserId); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "400"
		return res, nil
	} else {
		res.FollowerCount = int32(followerCount)
		res.FollowingCount = int32(followingCount)
	}

	return res, nil
}
