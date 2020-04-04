package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	ID    int             `db:"id"`
	Name  string          `db:"name"`
	Sex   sql.NullString  `db:"sex"`
	Score sql.NullFloat64 `db:"score"`
}

func loginDB(db string, loginInfo string) (*sql.DB, error) {
	DB, err := sql.Open(db, loginInfo)
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func mysqlSelect(DB *sql.DB) {
	sqlStr := "SELECT * FROM tb_stu where id>?"
	//row := DB.QueryRow(sqlStr, 1)  查询单行数据
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,error:%v\n", err)
	}
	defer func() { //注意查询占用连接池资源，要么被释放掉，要么手动关闭
		if rows != nil {
			rows.Close()
		}
	}()

	var stu student
	for rows.Next() {
		err := rows.Scan(&stu.ID, &stu.Name, &stu.Sex, &stu.Score)
		if err != nil {
			fmt.Printf("scan failed,error:%v\n", err)
		}
		fmt.Printf("%#v\n", stu)
	}
}

func mysqlPrepareSelect(DB *sql.DB) {
	sqlStr := "SELECT * FROM tb_stu where id=?"
	//row := DB.QueryRow(sqlStr, 1)  查询单行数据
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed,error:%v\n", err)
	}
	defer func() { //注意关闭预处理
		//查询占用连接池资源，要么被释放掉，要么手动关闭
		if stmt != nil {
			stmt.Close()
		}
	}()

	/*
		var row *sql.Row
		defer func() { //查询占用连接池资源，要么被释放掉，要么手动关闭
			if row != nil {
				row.Close()
			}
		}()*/

	var stu student
	for i := 0; i < 6; i++ {
		row := stmt.QueryRow(i)
		err := row.Scan(&stu.ID, &stu.Name, &stu.Sex, &stu.Score)
		if err != nil {
			fmt.Printf("scan failed,error:%v\n", err)
		}
		fmt.Printf("%#v %d\n", stu, i)
	}
}

func mysqlInsert(DB *sql.DB) {
	sqlStr := "INSERT INTO tb_stu(name, sex, score) VALUES(?,?,?)"
	result, err := DB.Exec(sqlStr, "Lily", "F", 81.5)
	if err != nil {
		fmt.Printf("execute failed,error:%v\n", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("can't get last insert id,error:%v\n", err)
	}
	fmt.Printf("last insert id:%d\n", id)
}

func mysqlUpdate(DB *sql.DB) {
	sqlStr := "UPDATE tb_stu SET sex=?,score=? WHERE name=?"
	result, err := DB.Exec(sqlStr, "F", 91.5, "Nancy")
	if err != nil {
		fmt.Printf("execute failed,error:%v\n", err)
	}
	amount, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("can't get last affect rows,error:%v\n", err)
	}
	fmt.Printf("%d rows affect\n", amount)
}

func mysqlDelete(DB *sql.DB) {
	sqlStr := "DELETE FROM tb_stu WHERE id=?"
	result, err := DB.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("execute failed,error:%v\n", err)
	}
	amount, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("can't get last affect rows,error:%v\n", err)
	}
	fmt.Printf("%d rows affect\n", amount)
}

func main() {
	DB, err := loginDB("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	//mysqlSelect(DB)
	//mysqlInsert(DB)
	//mysqlSelect(DB)
	//mysqlUpdate(DB)
	//mysqlSelect(DB)
	//mysqlDelete(DB)
	//mysqlSelect(DB)
	mysqlPrepareSelect(DB)
}
