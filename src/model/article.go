package model

import (
	"time"

	"github.com/go-xorm/xorm"

	"go-xn/src/util"
)

// Article model
type Article struct {
	ID         uint64    `xorm:"bigint(20) notnull autoincr pk 'ID'"`
	Title      string    `xorm:"text notnull 'title'"`
	Content    string    `xorm:"longtext notnull 'content'"`
	Views      uint64    `xorm:"bigint(20) notnull default(0) 'views'"`
	Status     int       `xorm:"tinyint(4) notnull default(0) 'status'"`
	CreateTime time.Time `xorm:"datetime created notnull default('0000-00-00 00:00:00') 'create_time'"`
	UpdateTime time.Time `xorm:"datetime updated notnull default('0000-00-00 00:00:00') 'update_time'"`
	DeleteTime time.Time `xorm:"datetime notnull default('0000-00-00 00:00:00') 'delete_time'"`
}

var orm *xorm.Engine = util.DBEngine()

// ArticlesExist if article exist
func ArticlesExist(id uint64) bool {
	has, _ := orm.Exist(&Article{
		ID: id,
	})

	return has
}

// ArticlesCount return count number
func ArticlesCount() int {
	article := &Article{}

	count, _ := orm.Count(article)

	return int(count)
}

// ArticleByID will return article by ID
func ArticleByID(id uint64) (*Article, bool) {
	article := &Article{
		ID: id,
	}

	has, _ := orm.Get(article)

	return article, has
}
