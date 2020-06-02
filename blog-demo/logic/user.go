package logic

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"blog-demo/dal/db"
	"blog-demo/model"
)

//GetUser 获取一个用户的全部信息
func GetUser(id uint64) (user *model.User, err error) {
	u := [1]uint64{id}
	users, err := db.SelectUser(u[:], "*")
	if err != nil {
		return
	}
	user = users[0]
	return
}

//CheckUser 校验用户密码并返回ID
func CheckUser(userName string, pw string) (id uint64, err error) {
	if len(userName) < 6 {
		err = fmt.Errorf("invaild user")
		return
	}

	pwR, id, err := db.GetUserPWoID(userName)
	if err != nil {
		return
	}

	if pw != pwR {
		err = fmt.Errorf("password error")
		return
	}

	return
}

//AddUser 添加一个用户，不能重复
func AddUser(userName string, pw string, pwCf string) error {
	if len(userName) < 6 {
		err := fmt.Errorf("username format incorrect")
		return err
	}
	if len(pw) < 6 {
		err := fmt.Errorf("password format incorrect")
		return err
	}
	if pw != pwCf {
		err := fmt.Errorf("111:two password is inconsistent")
		return err
	}

	rand.Seed(time.Now().UnixNano())

	u := &model.User{
		Name:     userName,
		Password: pw,
	}

	idS := time.Now().Format("20060102150405") + fmt.Sprintf("%d", rand.Uint32()%100000)
	id, err := strconv.ParseUint(idS, 10, 64)
	if err != nil {
		return err
	}
	u.ID = id

	err = db.InsertUser(u)
	if err != nil {
		return err
	}

	return nil
}
