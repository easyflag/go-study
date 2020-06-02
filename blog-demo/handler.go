package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"blog-demo/logic"
	"blog-demo/model"

	"github.com/gin-gonic/gin"
)

//LoginIDs 已登录用户ID表
var LoginIDs map[uint64]int

//IndexShow 首页显示
func IndexShow(c *gin.Context) {
	list, err := logic.OrganizeArticleListInfo(1, 8, "全部", "*")
	if err != nil {
		ErrorPrint(err, "Organize ArticleList Info failed", c)
		return
	}

	rank, err := logic.GetArticleListByViewCount()
	if err != nil {
		ErrorPrint(err, "get article ranking failed", c)
		return
	}

	categories, err := logic.GetCategoryList()
	if err != nil {
		ErrorPrint(err, "Get Category List failed", c)
		return
	}

	Render(c, gin.H{"articleList": list,
		"articleRank": rank,
		"categories":  categories,
		"logo":        "网站首页"}, "view/index.html")
}

//IndexPost 首页提交
func IndexPost(c *gin.Context) {
	if c.PostForm("submit") == "退出登录" {
		//清除客户端cookie
		//name, value string, maxAge int, path, domain string, secure, httpOnly bool
		c.SetCookie("token", "", -1, "", "", false, true)

		//删除已登录用户ID表的相应数据
		I, _ := c.Get("logInfo")
		delete(LoginIDs, I.(*model.LogInfo).UserID)
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

//RegisterShow 注册页面显示
func RegisterShow(c *gin.Context) {
	Render(c, gin.H{}, "view/register.html")
}

//RegisterPost 注册页面提交
func RegisterPost(c *gin.Context) {
	userName := c.PostForm("userName")
	userPwd := c.PostForm("userPwd")
	confirmPwd := c.PostForm("confirmPwd")

	err := logic.AddUser(userName, userPwd, confirmPwd)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "1062") {
			msg = "用户名已经存在，请重新输入."
		} else if strings.Contains(err.Error(), "111") {
			msg = "两次密码输入不一致，请重新输入."
		} else {
			msg = "请按规范填写信息."
		}

		Render(c, gin.H{"msg": msg}, "view/register.html")
		fmt.Println("add failed", err)
		return
	}

	//追加到已登录用户列表
	id, _ := logic.CheckUser(userName, userPwd)
	if LoginIDs == nil {
		LoginIDs = make(map[uint64]int)
	}
	LoginIDs[id]++

	//往客服端写cookie
	c.SetCookie("token", userName, 3600, "", "", false, true)

	fmt.Println(userName, userPwd, confirmPwd)

	c.Redirect(http.StatusMovedPermanently, "/")
}

//LoginShow 登录页面显示
func LoginShow(c *gin.Context) {
	Render(c, gin.H{}, "view/login.html")
}

//LoginPost 登录页面提交
func LoginPost(c *gin.Context) {
	userName := c.PostForm("userName")
	userPwd := c.PostForm("userPwd")

	id, err := logic.CheckUser(userName, userPwd)
	if err != nil {
		Render(c, gin.H{"msg": "用户名或密码错误.请注意区分大小写."}, "view/login.html")
		fmt.Println("pw err", err)
		return
	}

	//查重 防止一个账号重复登录
	if LoginIDs != nil {
		for i := range LoginIDs {
			if id == i {
				Render(c, gin.H{"msg": "该用户已在其它设备登录！"}, "view/login.html")
				fmt.Println("用户重复登录")
				return
			}
		}
	}

	fmt.Println("LoginSubmit", userName, userPwd)

	//追加到已登录用户列表
	if LoginIDs == nil {
		LoginIDs = make(map[uint64]int)
	}
	LoginIDs[id]++

	//往客服端写cookie
	c.SetCookie("token", userName, 3600, "", "", false, true)

	c.Redirect(http.StatusMovedPermanently, "/")
}

//ArticlePostShow 发表文章页面显示
func ArticlePostShow(c *gin.Context) {
	list, err := logic.GetCategoryList()
	if err != nil {
		fmt.Printf("Get CategoryList failed,error:%v\n", err)
		c.HTML(http.StatusOK, "view/500.html", nil)
		return
	}

	Render(c, gin.H{"categoryList": list,
		"logo": "发布文章"}, "view/submit_article.html")
}

//ArticlePostSubmit 发布文章页面提交
func ArticlePostSubmit(c *gin.Context) {
	list, err := logic.GetCategoryList()
	if err != nil {
		fmt.Printf("Organize ArticlePost data failed,error:%v\n", err)
		c.HTML(http.StatusOK, "view/500.html", nil)
		return
	}
	H := gin.H{"categoryList": list,
		"logo": "发布文章"}
	if I, _ := c.Get("logInfo"); I.(*model.LogInfo).Status == false {
		H["msg"] = "发布文章请先登录."
	} else {
		id := c.PostForm("categoryID")
		categoryID, _ := strconv.ParseUint(id, 10, 64)
		title := c.PostForm("title")
		content := c.PostForm("content")
		I, _ := c.Get("logInfo")
		err := logic.AddArticle(uint8(categoryID), I.(*model.LogInfo).UserID,
			title, content)
		if err != nil {
			fmt.Printf("add Article failed,error:%v\n", err)
			c.HTML(http.StatusOK, "view/500.html", nil)
			return
		}
		H["msg"] = "发布成功！"
	}

	Render(c, H, "view/submit_article.html")
}

//ArticleListShow 文章列表页显示
func ArticleListShow(c *gin.Context) {
	//获取url参数
	currentCategory := c.Query("category")
	if currentCategory == "" {
		currentCategory = "全部"
	}
	fmt.Println(currentCategory, "ArticleListShow")
	currentPageS := c.Query("page")
	if currentPageS == "" {
		currentPageS = "1"
	}
	currentPage, err := strconv.ParseUint(currentPageS, 10, 64)
	if err != nil {
		ErrorPrint(err, "param mistake", c)
		return
	}

	fmt.Println(currentPage, "ArticleListShow")

	list, err := logic.OrganizeArticleListInfo(uint(currentPage), 5, currentCategory, "*")
	if err != nil {
		ErrorPrint(err, "Organize ArticleList Info failed", c)
		return
	}

	s, _ := logic.GetAllArticlesID(currentCategory)
	total := len(s)
	fmt.Println(total, "ArticleListShow")
	var maxPage int
	if total%5 == 0 {
		maxPage = total / 5
	} else {
		maxPage = total/5 + 1
	}

	categories, err := logic.GetCategoryList()

	if err != nil {
		ErrorPrint(err, "Get Category List failed", c)
		return
	}

	Render(c, gin.H{"articleList": list,
		"currentCategory": currentCategory,
		"categories":      categories,
		"total":           total,
		"currentPage":     currentPage,
		"maxPage":         maxPage,
		"nextPage":        currentPage + 1,
		"prevPage":        currentPage - 1,
		"logo":            "文章列表"}, "view/list.html")
}

//ArticleListPost 文章列表页提交
func ArticleListPost(c *gin.Context) {
	category := c.PostForm("category")
	fmt.Println(category, "ArticleListPost")

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/articles/list?category=%s", category))
}

//ViewArticleShow 文章详情页显示
func ViewArticleShow(c *gin.Context) {
	idS := c.Param("id")
	id, _ := strconv.ParseUint(idS, 10, 64)

	//文章相关
	article, err := logic.GetArticle(id)
	if err != nil {
		ErrorPrint(err, "Get Article failed", c)
		return
	}
	p := strings.Split(article.Content, "\r")

	//评论相关
	comments, err := logic.GetComment(id)
	if err != nil {
		fmt.Println(err)
	}

	userStrs := make([]string, 0, 10)
	for _, v := range comments {
		u, err := logic.GetUser(v.UserID)
		if err != nil {
			ErrorPrint(err, "Get User failed", c)
			return
		}
		userStrs = append(userStrs, u.Name)
	}
	u, err := logic.GetUser(article.UserID)
	if err != nil {
		ErrorPrint(err, "Get User failed", c)
		return
	}

	err = logic.ViewCountAddOne(id)
	if err != nil {
		ErrorPrint(err, "View Count Add One failed", c)
		return
	}

	Render(c, gin.H{
		"logo":         "浏览文章",
		"article":      article,
		"paragraph":    p,
		"comments":     comments,
		"commentUsers": userStrs,
		"articleUser":  u.Name}, "view/view_article.html")
}

//ViewArticlePost 文章详情页提交
func ViewArticlePost(c *gin.Context) {
	content := c.PostForm("content")
	idS := c.Param("id")
	id, _ := strconv.ParseUint(idS, 10, 64)

	I, _ := c.Get("logInfo")
	logInfo := I.(*model.LogInfo)
	if logInfo.Status == true {
		err := logic.AddComment(logInfo.UserID, id, content)
		if err != nil {
			ErrorPrint(err, "Add Comment failed", c)
		}
	} else {
		err := logic.AddComment(2020010119191900002, id, content)
		if err != nil {
			ErrorPrint(err, "Add Comment failed", c)
		}
	}

	err := logic.CommentCountAddOne(id)
	if err != nil {
		ErrorPrint(err, "Comment Count Add One failed", c)
		return
	}

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/articles/view/%d", id))
}

//UserShow 用户页面展示
func UserShow(c *gin.Context) {
	userName := c.Param("id")
	I, _ := c.Get("logInfo")
	logInfo := I.(*model.LogInfo)

	list, err := logic.OrganizeArticleListInfo(1, 999999, "", userName)
	if err != nil {
		ErrorPrint(err, "OrganizeArticleListInfo", c)
		return
	}

	fmt.Println(userName, "用户页面展示")
	Render(c, gin.H{"logo": "用户",
		"userName": userName,
		"articles": list,
		"logInfo":  logInfo}, "view/user.html")
}

//UserPost 用户页提交
func UserPost(c *gin.Context) {
	idS := c.Param("id")
	arID := c.PostForm("artcileID")
	articleID, _ := strconv.ParseUint(arID, 10, 64)

	err := logic.DeleteArticle(articleID)
	if err != nil {
		fmt.Println(err, "DeleteArticle")
	}

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/users/view/%s", idS))
}

//ArticleModShow 文章修改页展示
func ArticleModShow(c *gin.Context) {
	idS := c.Param("id")
	id, _ := strconv.ParseUint(idS, 10, 64)
	I, _ := c.Get("logInfo")
	logInfo := I.(*model.LogInfo)

	list, err := logic.GetCategoryList()
	if err != nil {
		ErrorPrint(err, "Get Category List failed", c)
		return
	}

	a, err := logic.GetArticle(id)
	if err != nil {
		ErrorPrint(err, "Get Article  failed", c)
		return
	}

	//用户鉴权
	fmt.Println(a.UserID, logInfo.UserID, "用户鉴权")
	if a.UserID != logInfo.UserID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	Render(c, gin.H{"categoryList": list,
		"article": a,
		"logo":    "发布文章"}, "view/submit_article.html")
}

//ArticleModPost 文章修改页提交
func ArticleModPost(c *gin.Context) {
	idS := c.Param("id")
	articleID, _ := strconv.ParseUint(idS, 10, 64)
	categoryID := c.PostForm("categoryID")
	title := c.PostForm("title")
	content := c.PostForm("content")
	fmt.Println(idS, articleID, categoryID, title, content)

	err := logic.ModArticle(articleID, title, categoryID, content)
	if err != nil {
		fmt.Println(err, "ModArticle")
	}

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/articles/view/%d", articleID))
}
