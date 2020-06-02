package db

import (
	"testing"

	"blog-demo/model"
)

func init() {
	err := Init()
	if err != nil {
		panic(err)
	}
}

func TestInsertComment(t *testing.T) {
	c := new(model.Comment)
	c.ID = 2020010119193300005
	c.UserID = 2020031915312145321
	c.ArticleID = 2020032101572675803
	c.Content = "fdsafasdfasdfads"

	err := InsertComment(c)
	if err != nil {
		t.Errorf("insert comment failed,error:%v\n", err)
	}
}

func TestSelectComment(t *testing.T) {
	ids := [2]uint64{2020032101572675803, 2020031915314771905}
	comments, err := SelectComent(ids[:], "*")
	if err != nil {
		t.Errorf("get comments failed,error:%v\n", err)
	}

	for _, v := range comments {
		t.Logf("%#v\n", *v)
	}
}
