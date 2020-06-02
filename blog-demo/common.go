package controller

import (
	"fmt"
	"net/http"

	"blog-demo/model"

	"github.com/gin-gonic/gin"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {
	//从gin上下文中获取登录信息，这是每个页面都共用的
	I, _ := c.Get("logInfo")
	data["logInfo"] = I.(*model.LogInfo)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

//ErrorPrint 输出错误页面
func ErrorPrint(err error, str string, c *gin.Context) {
	fmt.Printf("%s,error:%v\n", str, err)
	c.HTML(http.StatusOK, "view/500.html", nil)
}
