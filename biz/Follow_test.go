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
