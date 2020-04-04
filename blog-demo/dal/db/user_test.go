package db

import (
	"testing"
)

func init() {
	err := Init()
	if err != nil {
		panic(err)
	}
}

func TestSelectUser(t *testing.T) {
	var IDs [2]uint64

	IDs[0] = 2020010119191900001
	IDs[1] = 2020010119191900002

	users, err := SelectUser(IDs[:], "*")
	if err != nil {
		t.Errorf("get user info failed,error:%v\n", err)
	}

	t.Logf("%#v\n", users[0])
}
