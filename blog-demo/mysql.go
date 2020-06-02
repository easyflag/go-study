package db

import (
	_ "github.com/go-sql-driver/mysql" //导入驱动
	"github.com/jmoiron/sqlx"
)

//DB 数据库实例
var DB *sqlx.DB

//Init 连接数据库
func Init() (err error) {
	DB, err = sqlx.Open("mysql",
		"root:123456@tcp(localhost:3306)/blog-demo?parseTime=true&loc=Local")
	if err != nil {
		return
	}
	DB.SetMaxIdleConns(16)
	DB.SetMaxOpenConns(100)
	err = DB.Ping()
	if err != nil {
		return
	}

	return nil
}
