package db

import (
	"testing"
	"time"

	"github.com/go-study/blog-demo/model"
)

func init() {
	err := Init()
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	a := new(model.Article)
	a.ID = 2020010119191900004
	a.Status = 1
	a.CategoryID = 3
	a.UserID = 2020010119191900002
	a.CreateTime = time.Now()
	a.UpdateTime = time.Now()
	a.ViewCount = 1
	a.CommentCount = 1
	a.Title = "test1"
	a.Content = "fdsd1515"
	a.Summary = "fdsd1515"

	err := InsertArticle(a)
	if err != nil {
		t.Errorf("insert article failed,error:%v\n", err)
	}
}

func TestGetArticleList(t *testing.T) {
	list, err := GetArticleListInCategory(1, 10, 0)
	if err != nil {
		t.Errorf("get article list failed,error:%v\n", err)
	}

	for _, v := range list {
		t.Logf("%#v\n", v)
	}
}

func TestSelectArticle(t *testing.T) {
	list, err := GetArticleListInCategory(1, 10, 0)
	if err != nil {
		t.Errorf("get article list failed,error:%v\n", err)
	}

	articles, err := SelectArticle(list, "view_count", "comment_count", "title", "summary")
	if err != nil {
		t.Errorf("get article list failed,error:%v\n", err)
	}

	for _, v := range articles {
		t.Logf("%#v\n", *v)
	}
}

func TestUpdateArticlet(t *testing.T) {
	KV := map[string]interface{}{"view_count": 11, "comment_count": 11}

	err := UpdateArticle(2020032916563057733, KV)
	if err != nil {
		t.Errorf("update article failed,error:%v\n", err)
	}
}
