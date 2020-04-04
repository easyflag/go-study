package logic

import (
	"testing"

	"github.com/go-study/blog-demo/dal/db"
)

func init() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
}

func TestOrganizeArticleListInfo(t *testing.T) {
	list, err := OrganizeArticleListInfo(1, 999999, "杂闻", "张三")
	if err != nil {
		t.Errorf("Organize Article List Info failed,error:%v\n", err)
	}

	t.Error(list[0].UpdateTime)
	t.Log(list)
}

func TestGetAllArticlesID(t *testing.T) {
	s, err := GetAllArticlesID("汽车")
	if err != nil {
		t.Errorf("Get Article failed,error:%v\n", err)
	}
	t.Log(s)
}

func TestGetArticle(t *testing.T) {
	a, err := GetArticle(2020032101572675803)
	if err != nil {
		t.Errorf("Get Article failed,error:%v\n", err)
	}
	t.Log(a)
}

func TestAddArticle(t *testing.T) {
	err := AddArticle(1, 2020031000520282260, "多福多寿", `而另一方面我们又推迟我们的满足感，
	推迟我们应优先考虑的事情，推迟我们的幸福感，常常说服自己“有朝一日”会比今天更好。不幸的是，
	如此告诫我们朝前看的大脑动力只能重复来重复去，以致“有朝一日”永远不会真正来临。`)

	if err != nil {
		t.Errorf("Add Article failed,error:%v\n", err)
	}
}

func TestGetComment(t *testing.T) {
	c, err := GetComment(202003210157267580)
	if err != nil {
		t.Errorf("Get Comment failed,error:%v\n", err)
	}
	t.Log(c)
}

func TestAddComment(t *testing.T) {
	err := AddComment(2020031915312145321, 2020032101572675803, "dfrf")

	if err != nil {
		t.Errorf("Add Comment failed,error:%v\n", err)
	}
}

func TestModArticle(t *testing.T) {
	err := ModArticle(2020032916563057733, "Windows 10用户有连不上网的情况，微软：正在解决",
		"6",
		`据外媒最新报道称，微软证实了一个新的Bug，那就是网的问题。`)
	if err != nil {
		t.Errorf("Mod Article failed,error:%v\n", err)
	}
}

func TestDeleteArticle(t *testing.T) {
	err := DeleteArticle(2020032916563057733)
	if err != nil {
		t.Errorf("Mod Article failed,error:%v\n", err)
	}
}

func TestViewCountAddOne(t *testing.T) {
	err := ViewCountAddOne(2020032101582878789)
	if err != nil {
		t.Errorf("View Count Add One failed,error:%v\n", err)
	}
}

func TestCommentCountAddOne(t *testing.T) {
	err := CommentCountAddOne(2020032101582878789)
	if err != nil {
		t.Errorf("View Count Add One failed,error:%v\n", err)
	}
}
