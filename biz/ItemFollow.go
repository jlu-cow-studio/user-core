package biz

import (
	"github.com/jlu-cow-studio/common/dal/mysql"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
)

func AddUserItemFollow(userId int, itemId int) error {
	return mysql.GetDBConn().Table("user_item_follow").Create(&struct{
		UserId int `gorm:"column:user_id"`
		ItemId int `gorm:"colume:item_id"`
	}{
		UserId: userId,
		ItemId: itemId,
	}).Error
}

func DelUserItemFollow(userId int, itemId int) error {
	return mysql.GetDBConn().Table("user_follow").Delete(nil, " user_id = ? and  item_id = ?", userId, itemId).Error
}

func UserItemFollowList(userId int) ([]common.Item, error) {
	itemIds := make([]int, 0)
	err := mysql.GetDBConn().Table("user_follow").Where("user_id = ?", userId).Find(&itemIds).Error
	return itemIds, err
}