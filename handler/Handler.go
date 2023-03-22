package handler

import "github.com/jlu-cow-studio/common/dal/rpc/user_core"

type Handler struct {
	user_core.UnimplementedUserCoreServiceServer
}
