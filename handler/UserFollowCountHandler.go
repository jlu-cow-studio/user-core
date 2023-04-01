package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
)

func (h *Handler) FollowCount(ctx context.Context, req *user_core.FollowCountReq) (res *user_core.FollowCountRes, err error) {

	return res, nil
}
