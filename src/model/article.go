package model

import "time"

// Article model
type Article struct {
	ID         uint64    `form:"article_id" xorm:"bigint(20) notnull autoincr"`
	Title      string    `form:"title" xorm:"text notnull"`
	Content    string    `form:"content" xorm:"longtext notnull"`
	Views      uint64    `form:"views" xorm:"bigint(20) notnull default(0)"`
	Status     int       `form:"status" xorm:"tinyint(1) notnull default(0)"`
	CreateTime time.Time `form:"create_time" xorm:"datetime created notnull"`
	UpdateTime time.Time `form:"update_time" xorm:"datetime updated notnull"`
	DeleteTime time.Time `form:"delete_time" xorm:"datetime deleted notnull"`
}
