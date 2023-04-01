package biz

import (
	"github.com/jlu-cow-studio/common/dal/mysql"
	"github.com/jlu-cow-studio/common/dal/rpc/user_core"
)

func CheckUserFollowed(followerId, followingId string) (bool, error) {

	count := new(int64)
	tx := mysql.GetDBConn().Table("user_follow").Where("follower_id = ?", followerId).Where("following_id = ?", followingId).Count(count)

	return *count == 1, tx.Error
}

func AddUserFollow(followerId, followingId string) error {

	return mysql.GetDBConn().Table("user_follow").Create(&struct {
		FollowerId  string `gorm:"column:follower_id"`
		FollowingId string `gorm:"column:following_id"`
	}{
		FollowerId:  followerId,
		FollowingId: followingId,
	}).Error

}

func DelUserFollow(followerId, followingId string) error {
	return mysql.GetDBConn().Table("user_follow").Delete(nil, " follower_id = ? and  following_id = ?", followerId, followingId).Error
}

func FollowingList(followerId string, offset, size int) ([]*user_core.UserInfo, error) {
	list := []*user_core.UserInfo{}
	tx := mysql.GetDBConn().Table("user_follow").
		Joins("left join user on user_follow.following_id = user.uid").Where("user_follow.follower_id = ?", followerId).
		Select("user.*").Offset(offset).Limit(size).Find(&list)
	return list, tx.Error
}

func FollowerList(followingId string, offset, size int) ([]*user_core.UserInfo, error) {
	list := []*user_core.UserInfo{}
	tx := mysql.GetDBConn().Table("user_follow").
		Joins("left join user on user_follow.follower_id = user.uid").Where("user_follow.following_id = ?", followingId).
		Select("user.*").Offset(offset).Limit(size).Find(&list)
	return list, tx.Error
}

func FollowCount(userId string) (int, int, error) {

	count := &struct {
		FollowerCount  int `gorm:"column:follower_count"`
		FollowingCount int `gorm:"column:following_count"`
	}{}

	tx := mysql.GetDBConn().Table("user_follow_count").Find(count, "uid = ?", userId)

	return count.FollowerCount, count.FollowingCount, tx.Error
}
