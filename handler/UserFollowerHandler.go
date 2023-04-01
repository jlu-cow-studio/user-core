package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
)

func (h *Handler) Follower(ctx context.Context, req *user_core.FollowersReq) (res *user_core.FollowersRes, err error) {

	return res, nil
}
