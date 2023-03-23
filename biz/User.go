package biz

import (
	"github.com/jlu-cow-studio/common/dal/mysql"
	dao "github.com/jlu-cow-studio/common/model/dao_struct/mysql"
)

func CheckUserExsit(username string) (bool, error) {

	conn := mysql.GetDBConn()
	var count int64

	if err := conn.Table("user").Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}

	return count >= 1, nil
}

func GetUserInfo(username string) (dao.User, error) {

	user := &dao.User{}
	conn := mysql.GetDBConn()

	if err := conn.Table("user").Where("username = ?", username).First(&user).Error; err != nil {
		return *user, err
	}

	return *user, nil
}
