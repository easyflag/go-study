package model

import "time"

/*业务逻辑数据结构*/

//ArticleDetail 文章展示列表的整合信息
type ArticleDetail struct {
	CategoryName string    `db:"class"`         //文章分类
	UserName     string    `db:"name"`          //用户名
	UpdateTime   time.Time `db:"update_time"`   //文章更新时间
	ViewCount    uint32    `db:"view_count"`    //阅读数
	CommentCount uint32    `db:"comment_count"` //评论数
	ID           uint64    `db:"id"`            //文章ID
	Title        string    `db:"title"`         //文章标题
	Summary      string    `db:"summary"`       // 文章摘要
}

//LogInfo 登录信息
type LogInfo struct {
	Status   bool
	UserID   uint64
	UserName string
}
