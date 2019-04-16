package model

import (
	"time"

	"github.com/go-xorm/xorm"

	"go-xn/src/util"
)

// Article model
type Article struct {
	ID         uint64    `json:"ID" xorm:"bigint(20) notnull autoincr pk 'ID'"`
	Title      string    `json:"title" xorm:"text notnull 'title'"`
	Content    string    `json:"content" xorm:"longtext notnull 'content'"`
	Views      uint64    `json:"article_views" xorm:"bigint(20) notnull default(0) 'article_views'"`
	Status     int       `json:"article_status" xorm:"tinyint(4) notnull default(0) 'article_status'"`
	CreateTime time.Time `json:"create_time" xorm:"datetime created notnull default('0000-00-00 00:00:00') 'create_time'"`
	UpdateTime time.Time `json:"update_time" xorm:"datetime updated notnull default('0000-00-00 00:00:00') 'update_time'"`
	DeleteTime time.Time `json:"delete_time" xorm:"datetime notnull default('0000-00-00 00:00:00') 'delete_time'"`
}

var orm *xorm.Engine = util.DBEngine()

// HomeView return articles while index page
func HomeView() ([]Article, error) {
	db := orm.NewSession()
	defer db.Close()

	articles := make([]Article, 0, 4)

	err := db.Table("article").
		Cols("ID", "title", "content", "article_views").
		Where("article_status = 1").
		Desc("create_time").
		Find(&articles)

	return articles, err
}

// ArticlesExist if article exist
func ArticlesExist(id uint64) bool {
	db := orm.NewSession()
	defer db.Close()

	has, _ := db.Exist(&Article{
		ID: id,
	})

	return has
}

// ArticlesCount return count number
func ArticlesCount() int {
	db := orm.NewSession()
	defer db.Close()

	article := &Article{}

	count, _ := db.Count(article)

	return int(count)
}

// ArticleByID will return article by ID
func ArticleByID(id uint64) (*Article, bool) {
	db := orm.NewSession()
	defer db.Close()

	article := &Article{
		ID: id,
	}

	has, _ := db.Get(article)

	return article, has
}
