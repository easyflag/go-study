package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/blog-demo/controller"
	"github.com/go-study/blog-demo/dal/db"
)

var router *gin.Engine

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}

	router = gin.Default()

	RoutesInit()

	router.Run(":88")
}

//RoutesInit 初始化所有路径
func RoutesInit() {
	//加载页面模板与静态数据
	router.LoadHTMLGlob("view/*")
	router.Static("/static/", "./static")

	//启用全局中间件
	router.Use(controller.SetLogInfo)

	//首页
	router.GET("/", controller.IndexShow)
	router.POST("/", controller.IndexPost)

	//用户组
	usersRoute := router.Group("/users")
	{
		usersRoute.GET("/login", controller.EnsureNotLogin, controller.LoginShow)
		usersRoute.POST("/login", controller.LoginPost)
		usersRoute.GET("/register", controller.EnsureNotLogin, controller.RegisterShow)
		usersRoute.POST("/register", controller.RegisterPost)
		usersRoute.GET("/view/:id", controller.UserShow)
		usersRoute.POST("/view/:id", controller.EnsureLogin, controller.UserPost)
	}

	//文章组
	articlesRoute := router.Group("/articles")
	{
		articlesRoute.GET("/submit", controller.ArticlePostShow)
		articlesRoute.POST("/submit", controller.ArticlePostSubmit)
		articlesRoute.GET("/mod/:id", controller.EnsureLogin, controller.ArticleModShow)
		articlesRoute.POST("/mod/:id", controller.EnsureLogin, controller.ArticleModPost)
		articlesRoute.GET("/list", controller.ArticleListShow)
		articlesRoute.POST("/list", controller.ArticleListPost)
		articlesRoute.GET("/view/:id", controller.ViewArticleShow)
		articlesRoute.POST("/view/:id", controller.ViewArticlePost)
	}
}
