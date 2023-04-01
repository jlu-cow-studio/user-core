package biz

import "github.com/jlu-cow-studio/common/dal/mysql"

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
