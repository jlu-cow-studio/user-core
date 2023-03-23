package biz

import (
	"errors"

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

func InsertUser(user *dao.User) error {
	conn := mysql.GetDBConn()
	return conn.Table("user").Create(user).Error
}

func CheckUserAuthIsNot(uid, role string) error {
	conn := mysql.GetDBConn()

	user := dao.User{}
	if err := conn.Table("user").Where("uid = ?", uid).First(&user).Error; err != nil {
		return err
	}

	if user.Role != role {
		return errors.New("User has been authorized!")
	}

	return nil
}

func UpdateUserRole(uid string, role string) error {
	conn := mysql.GetDBConn()

	if err := conn.Table("user").Where("uid = ?", uid).Update("role", role).Error; err != nil {
		return err
	}

	return nil
}
