package model

import "time"

/*与数据库一一对应的数据结构*/

//User 用户（与数据库对应）
type User struct {
	ID           uint64    `db:"id"`
	Status       int8      `db:"status"`
	Name         string    `db:"name"`
	RegisterTime time.Time `db:"register_time"`
	Password     string    `db:"password"`
}

//Article 文章（与数据库对应）
type Article struct {
	ID           uint64    `db:"id"`            //文章编号
	Status       int8      `db:"status"`        //文章状态
	CategoryID   uint8     `db:"category_id"`   //文章分类编号
	UserID       uint64    `db:"user_id"`       //用户编号
	CreateTime   time.Time `db:"create_time"`   //文章创建时间
	UpdateTime   time.Time `db:"update_time"`   //文章更新时间
	ViewCount    uint32    `db:"view_count"`    //阅读数
	CommentCount uint32    `db:"comment_count"` //评论数
	Title        string    `db:"title"`         //文章标题
	Summary      string    `db:"summary"`       // 文章摘要
	Content      string    `db:"content"`       //文章内容
}

//Category 分类（与数据库对应）
type Category struct {
	ID    uint8  `db:"id"`
	Class string `db:"class"`
}

//Comment 评论（与数据库对应）
type Comment struct {
	ID         uint64    `db:"id"`
	Status     int8      `db:"status"`
	UserID     uint64    `db:"user_id"`
	ArticleID  uint64    `db:"article_id"`
	CreateTime time.Time `db:"create_time"`
	Content    string    `db:"content"`
}
