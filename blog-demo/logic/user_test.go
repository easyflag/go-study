package logic

import (
	"testing"

	"github.com/go-study/blog-demo/dal/db"
)

func init() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
}

func TestGetUser(t *testing.T) {
	u, err := GetUser(2020031000520282260)
	if err != nil {
		t.Errorf("Get Article failed,error:%v\n", err)
	}
	t.Log(u)
}

func TestAddUser(t *testing.T) {
	err := AddUser("vistor", "123", "123")
	if err != nil {
		t.Errorf("Add User failed,error:%v\n", err)
	}

}

func TestCheckUser(t *testing.T) {
	_, err := CheckUser("vistor", "12p3")
	if err != nil {
		t.Errorf("Add User failed,error:%v\n", err)
	}

}
