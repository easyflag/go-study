package logic

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"blog-demo/dal/db"

	"blog-demo/model"
)

//OrganizeArticleListInfo 把文章展示列表需要的信息整合在一起
func OrganizeArticleListInfo(pageNum, pageSize uint, category, userName string) (articlelist []*model.ArticleDetail, err error) {
	var IDs []uint64
	if userName == "*" {
		id, err1 := db.GetCategoryIDByName(category)
		if err1 != nil {
			err = err1
			return
		}

		IDs, err = db.GetArticleListInCategory(pageNum, pageSize, id)
		if err != nil || IDs == nil {
			return
		}
	} else {
		_, id, err1 := db.GetUserPWoID(userName)
		if err1 != nil {
			err = err1
			return
		}

		IDs, err = db.GetArticlesInUser(id)
		if err != nil || IDs == nil {
			return
		}
	}

	articles, err := db.SelectArticle(IDs, "id",
		"user_id", "category_id", "update_time", "view_count", "comment_count", "title", "summary")
	if err != nil {
		return
	}

	userids := make([]uint64, 0, len(articles))
	categoryids := make([]uint8, 0, len(articles))
	for _, v := range articles {
		userids = append(userids, v.UserID)
		categoryids = append(categoryids, v.CategoryID)
	}

	users, err := db.SelectUser(userids, "name")
	if err != nil {
		return
	}

	categories, err := db.SelectCategory(categoryids, "class")
	if err != nil {
		return
	}

	articlelist = make([]*model.ArticleDetail, len(articles))

	for i := 0; i < len(articles); i++ {
		articlelist[i] = new(model.ArticleDetail)

		articlelist[i].CategoryName = categories[i].Class
		articlelist[i].UserName = users[i].Name
		articlelist[i].UpdateTime = articles[i].UpdateTime
		articlelist[i].ViewCount = articles[i].ViewCount
		articlelist[i].CommentCount = articles[i].CommentCount
		articlelist[i].ID = articles[i].ID
		articlelist[i].Title = articles[i].Title
		articlelist[i].Summary = articles[i].Summary
	}

	return
}

//GetArticleListByViewCount 获取文章阅读排行列表
func GetArticleListByViewCount() (articles []*model.Article, err error) {
	IDs, err := db.GetArticleListByViewCount(1, 15)
	if err != nil {
		return
	}

	articles, err = db.SelectArticle(IDs, "id", "title")
	if err != nil {
		return
	}

	return
}

//GetAllArticlesID 获取一个分类下的所有文章的ID
func GetAllArticlesID(name string) (IDs []uint64, err error) {
	id, err := db.GetCategoryIDByName(name)
	if err != nil {
		return
	}

	IDs, err = db.GetALLArticlesInCategory(id)
	if err != nil {
		return
	}

	return
}

//GetArticle 获取一篇文章的全部信息
func GetArticle(id uint64) (article *model.Article, err error) {
	a := [1]uint64{id}
	articles, err := db.SelectArticle(a[:], "*")
	if err != nil {
		return
	}
	article = articles[0]
	return
}

//GetComment 获取一则文章评论的全部信息
func GetComment(articleID uint64) (comments []*model.Comment, err error) {
	c := [1]uint64{articleID}
	comments, err = db.SelectComent(c[:], "*")
	if err != nil {
		return
	}
	return
}

//GetCategoryList 获取分类名
func GetCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetCategoryList()
	if err != nil {
		return
	}
	return
}

//AddArticle 向数据库添加文章
func AddArticle(categoryID uint8, userID uint64, title string, content string) error {
	rand.Seed(time.Now().UnixNano())

	a := &model.Article{
		CategoryID: categoryID,
		UserID:     userID,
		Title:      title,
		Content:    content,
	}

	idS := time.Now().Format("20060102150405") + fmt.Sprintf("%d", rand.Uint32()%100000)
	id, err := strconv.ParseUint(idS, 10, 64)
	if err != nil {
		return err
	}
	a.ID = id

	contentUTF := []rune(content)
	var i int
	var f int
	for i = 0; i < len(contentUTF); i++ {
		if contentUTF[i] == '.' || contentUTF[i] == '。' || contentUTF[i] == '!' || contentUTF[i] == '！' || contentUTF[i] == '？' || contentUTF[i] == '?' {
			f++
		} else if (contentUTF[i] == ',' || contentUTF[i] == '，') && f > 1 {
			break
		}
	}
	if f > 1 {
		a.Summary = string(contentUTF[:i+2]) + "..."
	} else {
		a.Summary = content
	}

	err = db.InsertArticle(a)
	if err != nil {
		return err
	}

	return nil
}

//AddComment 向数据库添加评论
func AddComment(userID, articleID uint64, content string) error {
	if content == "" {
		return fmt.Errorf("comment is empty")
	}

	rand.Seed(time.Now().UnixNano())

	c := &model.Comment{
		UserID:    userID,
		ArticleID: articleID,
		Content:   content,
	}

	idS := time.Now().Format("20060102150405") + fmt.Sprintf("%d", rand.Uint32()%100000)
	id, err := strconv.ParseUint(idS, 10, 64)
	if err != nil {
		return err
	}
	c.ID = id

	err = db.InsertComment(c)
	if err != nil {
		return err
	}

	return nil
}

//ModArticle 修改一篇文章，包括标题、分类、内容、摘要、更新时间
func ModArticle(id uint64, title, cateID, content string) error {
	//重建摘要
	contentUTF := []rune(content)
	var i, f int
	var summary string
	for i = 0; i < len(contentUTF); i++ {
		if contentUTF[i] == '.' || contentUTF[i] == '。' || contentUTF[i] == '!' || contentUTF[i] == '！' || contentUTF[i] == '？' || contentUTF[i] == '?' {
			f++
		} else if (contentUTF[i] == ',' || contentUTF[i] == '，') && f > 1 {
			break
		}
	}
	if f > 1 {
		summary = string(contentUTF[:i+2]) + "..."
	} else {
		summary = content
	}

	//更新时间
	updateTime := time.Now() //.Format("2006-01-02 15:04:05")

	err := db.UpdateArticle(id, map[string]interface{}{
		"title":       title,
		"category_id": cateID,
		"content":     content,
		"summary":     summary,
		"update_time": updateTime})
	if err != nil {
		return err
	}

	return nil
}

//DeleteArticle 删除一篇文章
func DeleteArticle(id uint64) error {
	err := db.UpdateArticle(id, map[string]interface{}{"status": -1})
	if err != nil {
		return err
	}

	return nil
}

//ViewCountAddOne 文章浏览数+1
func ViewCountAddOne(id uint64) error {
	articles, err := db.SelectArticle([]uint64{id}, "view_count")
	if err != nil {
		return err
	}
	n := fmt.Sprintf("%d", articles[0].ViewCount+1)

	err = db.UpdateArticle(id, map[string]interface{}{"view_count": n})
	if err != nil {
		return err
	}

	return nil
}

//CommentCountAddOne 文章评论数+1
func CommentCountAddOne(id uint64) error {
	articles, err := db.SelectArticle([]uint64{id}, "comment_count")
	if err != nil {
		return err
	}
	n := fmt.Sprintf("%d", articles[0].CommentCount+1)

	err = db.UpdateArticle(id, map[string]interface{}{"comment_count": n})
	if err != nil {
		return err
	}

	return nil
}
