package model

import (
	"time"

	"github.com/go-xorm/xorm"

	"go-xn/src/util"
)

// Article model
type Article struct {
	ID         uint64    `xorm:"bigint(20) notnull autoincr pk 'ID'"`
	Title      string    `xorm:"text notnull 'article_title'"`
	Content    string    `xorm:"longtext notnull 'article_content'"`
	Views      uint64    `xorm:"bigint(20) notnull default(0) 'article_views'"`
	Status     int       `xorm:"tinyint(4) notnull default(0) article_status"`
	CreateTime time.Time `xorm:"datetime created notnull default('0000-00-00 00:00:00') 'create_time'"`
	UpdateTime time.Time `xorm:"datetime updated notnull default('0000-00-00 00:00:00') 'update_time'"`
	DeleteTime time.Time `xorm:"datetime deleted notnull default('0000-00-00 00:00:00') 'delete_time'"`
}

var orm *xorm.Engine = util.DBEngine()

// ArticlesExist if article exist
func ArticlesExist(id uint64) bool {
	has, _ := orm.Exist(&Article{
		ID: id,
	})

	return has
}
