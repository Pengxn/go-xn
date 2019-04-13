package model

import (
	"time"

	"github.com/go-xorm/xorm"

	"go-xn/src/util"
)

// Article model
type Article struct {
	ID         uint64    `xorm:"bigint(20) notnull autoincr pk 'article_id'"`
	Title      string    `xorm:"text notnull 'article_title'"`
	Content    string    `xorm:"longtext notnull 'article_content'"`
	Views      uint64    `xorm:"bigint(20) notnull default(0) 'article_views'"`
	Status     int       `xorm:"tinyint(1) notnull default(0) article_status"`
	CreateTime time.Time `xorm:"datetime created notnull 'create_time'"`
	UpdateTime time.Time `xorm:"datetime updated notnull 'update_time'"`
	DeleteTime time.Time `xorm:"datetime deleted notnull 'delete_time'"`
}

var orm *xorm.Engine = util.DBEngine()

// GetArticles get all articles
func (a *Article) GetArticles() error {
	_, err := orm.Id(a.ID).Get(a)

	return err
}

// Exist if article exist
func (a *Article) Exist() (bool, error) {
	return orm.Get(a)
}
