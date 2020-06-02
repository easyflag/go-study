package db

import (
	"fmt"

	"blog-demo/model"
)

//SelectComent 查询文章ID，根据匹配的字段名，从数据库获取评论信息
func SelectComent(ArticleIDs []uint64, fieldName ...string) (comments []*model.Comment, err error) {
	var sqlS string
	for i, v := range fieldName {
		sqlS += v
		if i < len(fieldName)-1 {
			sqlS += ","
		}
	}

	sqlS = "SELECT " + sqlS + " FROM comments WHERE article_id=? ORDER BY create_time DESC"

	for _, v := range ArticleIDs {
		err = DB.Select(&comments, sqlS, v)
		if err != nil {
			return
		}
	}

	if comments == nil {
		err = fmt.Errorf("no such id")
		return
	}

	return
}

//InsertComment 向数据库插入评论
func InsertComment(c *model.Comment) error {
	if c == nil {
		err := fmt.Errorf("invalid comment parameter")
		return err
	}

	sqlS := `INSERT INTO comments(
		id, user_id, article_id, content) 
		  VALUES(?,?,?,?)`

	result, err := DB.Exec(sqlS, c.ID, c.UserID, c.ArticleID, c.Content)
	if err != nil {
		fmt.Printf("insert comment failed,error:%v\n", err)
		return err
	}

	id, _ := result.LastInsertId()
	fmt.Printf("last insert id is %d\n", id)

	return nil
}
