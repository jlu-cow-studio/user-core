package biz

import (
	"fmt"
	"testing"

	"github.com/jlu-cow-studio/common/dal/mysql"
)

func TestDelUserFollow(t *testing.T) {

	mysql.Init()

	fmt.Println(DelUserFollow("182", "322"))
}

func TestAddUserFollow(t *testing.T) {
	mysql.Init()

	fmt.Println(AddUserFollow("182", "322"))
}

func TestFollowingList(t *testing.T) {
	mysql.Init()

	fmt.Println(FollowingList("322", 0, 2))
}

func TestFollowerList(t *testing.T) {
	mysql.Init()

	fmt.Println(FollowerList("322", 0, 2))
}

func TestFollowCount(t *testing.T) {
	mysql.Init()

	fmt.Println(FollowCount("322"))
}
