package db

import (
	"fmt"

	"blog-demo/model"
)

//GetCategoryList 获取文章分类列表
func GetCategoryList() (categoryList []*model.Category, err error) {
	sqlS := `SELECT *
			FROM categories
			ORDER BY id`

	err = DB.Select(&categoryList, sqlS)
	if err != nil {
		return
	}

	return
}

//GetCategoryIDByName 通过分类名获取ID
func GetCategoryIDByName(name string) (ID uint8, err error) {
	sqlS := `SELECT id
			FROM categories
			WHERE class=?`

	row := DB.QueryRow(sqlS, name)
	if row == nil {
		err = fmt.Errorf("no such row in database")
		return
	}
	err = row.Scan(&ID)
	if err != nil {
		return
	}

	return
}

//SelectCategory 查询分类ID，根据匹配的字段名，从数据库获取分类信息
func SelectCategory(IDs []uint8, fieldName ...string) (categories []*model.Category, err error) {
	var sqlS string
	for i, v := range fieldName {
		sqlS += v
		if i < len(fieldName)-1 {
			sqlS += ","
		}
	}

	sqlS = "SELECT " + sqlS + " FROM categories WHERE id=?"

	for _, v := range IDs {
		err = DB.Select(&categories, sqlS, v)
		if err != nil {
			return
		}
	}

	if categories == nil {
		err = fmt.Errorf("no such id")
		return
	}

	return
}
