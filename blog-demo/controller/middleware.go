package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-study/blog-demo/dal/db"
	"github.com/go-study/blog-demo/model"
)

//SetLogInfo 获取客户端cookie，维持登录信息
func SetLogInfo(c *gin.Context) {
	if val, err := c.Cookie("token"); val != "" || err == nil {
		_, id, err := db.GetUserPWoID(val)
		if err != nil {
			fmt.Println("set log info failed,error: token error")
		}

		c.Set("logInfo", &model.LogInfo{
			Status:   true,
			UserName: val,
			UserID:   id,
		})
	} else {
		c.Set("logInfo", &model.LogInfo{
			Status:   false,
			UserName: "",
			UserID:   0,
		})
	}
}

//EnsureNotLogin 确保无登录状态
func EnsureNotLogin(c *gin.Context) {
	I, _ := c.Get("logInfo")
	if I.(*model.LogInfo).Status {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

//EnsureLogin 确保已经登录状态
func EnsureLogin(c *gin.Context) {
	I, _ := c.Get("logInfo")
	if !I.(*model.LogInfo).Status {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
