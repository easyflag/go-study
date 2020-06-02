package db

import (
	"fmt"

	"blog-demo/model"
)

//InsertArticle 向数据库插入文章
func InsertArticle(a *model.Article) error {
	if a == nil {
		err := fmt.Errorf("invalid article parameter")
		return err
	}

	sqlS := `INSERT INTO articles(
		id, category_id, user_id, title, content, summary) 
		  VALUES(?,?,?,?,?,?)`

	result, err := DB.Exec(sqlS, a.ID, a.CategoryID, a.UserID,
		a.Title, a.Content, a.Summary)
	if err != nil {
		fmt.Printf("insert article failed,error:%v\n", err)
		return err
	}

	id, _ := result.LastInsertId()
	fmt.Printf("last insert id is %d\n", id)

	return nil
}

//GetArticleListByViewCount 获取文章ID列表,以阅读数排序,限制获取数量
func GetArticleListByViewCount(pageNum, pageSize uint) (IDs []uint64, err error) {
	sqlS := `SELECT id
			FROM articles
			WHERE status=1
			ORDER BY view_count DESC
			LIMIT ?,?`

	err = DB.Select(&IDs, sqlS, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		return
	}

	return
}

//GetArticleListInCategory 获取指定分类的文章ID列表,以时间排序，限制获取数量
func GetArticleListInCategory(pageNum, pageSize uint, cateID uint8) (IDs []uint64, err error) {
	sqlS := ""
	if cateID != 0 {
		sqlS = `SELECT id
			FROM articles
			WHERE status=1 AND category_id=?
			ORDER BY view_count DESC
			LIMIT ?,?`
	} else {
		sqlS = `SELECT id
			FROM articles
			WHERE status=1 AND category_id>?
			ORDER BY update_time DESC
			LIMIT ?,?`
	}

	err = DB.Select(&IDs, sqlS, cateID, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		return
	}

	return
}

//GetALLArticlesInCategory 获取指定分类的文章ID列表,以时间排序，获取全部
func GetALLArticlesInCategory(cateID uint8) (IDs []uint64, err error) {
	sqlS := ""
	if cateID != 0 {
		sqlS = `SELECT id
			FROM articles
			WHERE status=1 AND category_id=?
			ORDER BY view_count DESC`
	} else {
		sqlS = `SELECT id
			FROM articles
			WHERE status=1 AND category_id>?
			ORDER BY update_time DESC`
	}

	err = DB.Select(&IDs, sqlS, cateID)
	if err != nil {
		return
	}

	return
}

//GetArticlesInUser 获取指定用户的文章ID列表,以时间排序，获取全部
func GetArticlesInUser(userID uint64) (IDs []uint64, err error) {
	sqlS := `SELECT id
			FROM articles
			WHERE status=1 AND user_id=?
			ORDER BY update_time DESC`

	err = DB.Select(&IDs, sqlS, userID)
	if err != nil {
		return
	}

	return
}

//SelectArticle 查询文章ID，根据匹配的字段名，从数据库获取文章信息
func SelectArticle(IDs []uint64, fieldNames ...string) (articles []*model.Article, err error) {
	var sqlS string
	for i, v := range fieldNames {
		sqlS += v
		if i < len(fieldNames)-1 {
			sqlS += ","
		}
	}

	sqlS = "SELECT " + sqlS + " FROM articles WHERE id=?"

	for _, v := range IDs {
		err = DB.Select(&articles, sqlS, v)
		if err != nil {
			return
		}
	}

	if articles == nil {
		err = fmt.Errorf("no such id")
		return
	}

	return
}

//UpdateArticle 从数据库中更新文章
func UpdateArticle(ID uint64, KV map[string]interface{}) error {
	var sqlS string
	for k := range KV {
		sqlS += k + "=?,"
	}

	sqlS = "UPDATE articles SET " + sqlS[:len(sqlS)-1] + " WHERE id=?"

	val := make([]interface{}, 0, len(KV)+1)
	for _, v := range KV {
		val = append(val, v)
	}
	val = append(val, ID)

	result, err := DB.Exec(sqlS, val...)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	fmt.Printf("last insert id is %d\n", id)

	return nil
}
