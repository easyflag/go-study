package db

import (
	"fmt"

	"github.com/go-study/blog-demo/model"
)

//SelectUser 查询用户ID，根据匹配的字段名，从数据库获取用户信息
func SelectUser(IDs []uint64, fieldName ...string) (users []*model.User, err error) {
	var sqlS string
	for i, v := range fieldName {
		sqlS += v
		if i < len(fieldName)-1 {
			sqlS += ","
		}
	}

	sqlS = "SELECT " + sqlS + " FROM users WHERE id=?"

	for _, v := range IDs {
		err = DB.Select(&users, sqlS, v)
		if err != nil {
			return
		}
	}

	if users == nil {
		err = fmt.Errorf("no such id")
		return
	}

	return
}

//InsertUser 向数据库插入用户
func InsertUser(u *model.User) error {
	if u == nil {
		err := fmt.Errorf("invalid user parameter")
		return err
	}

	sqlS := `INSERT INTO users(id,name,password) 
		  VALUES(?,?,?)`

	result, err := DB.Exec(sqlS, u.ID, u.Name, u.Password)
	if err != nil {
		fmt.Printf("insert article failed,error:%v\n", err)
		return err
	}

	id, _ := result.LastInsertId()
	fmt.Printf("last insert id is %d\n", id)

	return nil
}

//GetUserPWoID 由用户名查询用户密码和ID
func GetUserPWoID(userName string) (psw string, id uint64, err error) {
	sqlS := `SELECT password,id FROM users WHERE name=?`

	row := DB.QueryRow(sqlS, userName)
	err = row.Scan(&psw, &id)
	if err != nil {
		return
	}

	return
}


