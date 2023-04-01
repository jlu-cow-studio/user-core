package handler

import (
	"context"
	"math"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/user-core/biz"
)

func (h *Handler) Following(ctx context.Context, req *user_core.FollowingReq) (res *user_core.FollowingRes, err error) {

	res = &user_core.FollowingRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "498",
		},
	}

	offset := req.Page * req.PageSize
	limit := req.PageSize

	if req.Page < 0 || req.PageSize <= 0 {
		res.Base.Message = "invalid pagination"
		res.Base.Code = "402"
		return res, nil
	}

	if list, err := biz.FollowingList(req.UserId, int(offset), int(limit)); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "400"
		return res, nil
	} else if _, followingCount, err := biz.FollowCount(req.UserId); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "401"
		return res, nil
	} else {
		res.TotalCount = int32(followingCount)
		res.TotalPage = int32(math.Ceil(float64(followingCount) / float64(req.PageSize)))
		res.Users = list
	}

	return
}
