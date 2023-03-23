package handler

import (
	"context"

	"github.com/jlu-cow-studio/common/dal/rpc/base"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
	"github.com/jlu-cow-studio/common/model/dao_struct/mysql"
	mysql_model "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
	"github.com/jlu-cow-studio/user-core/biz"
)

func (h *Handler) UserRegister(ctx context.Context, req *user_core.UserRegisterReq) (res *user_core.UserRegisterRes, err error) {

	res = &user_core.UserRegisterRes{
		Base: &base.BaseRes{
			Message: "",
			Code:    "499",
		},
	}

	if ok, err := biz.CheckUserExsit(req.UserInfo.Username); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "401"
		return res, nil
	} else if ok {
		res.Base.Message = "User Already Exsit!"
		res.Base.Code = "402"
		return res, nil
	}

	user := &mysql.User{
		Username: req.UserInfo.Username,
		Password: req.UserInfo.Password,
		Role:     mysql_model.RoleNormal,
		Province: req.UserInfo.Province,
		City:     req.UserInfo.City,
		District: req.UserInfo.District,
	}

	if err := biz.InsertUser(user); err != nil {
		res.Base.Message = err.Error()
		res.Base.Code = "403"
		return res, nil
	}

	res.Base.Code = "200"
	return res, nil
}
