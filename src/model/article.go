package model

import (
	"time"

	"github.com/Pengxn/go-xn/src/util/log"
)

// Article model
type Article struct {
	ID         uint64     `json:"ID" xorm:"bigint(20) notnull autoincr pk 'ID'"`
	Slug       string     `json:"slug,omitempty" xorm:"varchar(255) unique notnull 'slug'"`
	Title      string     `json:"title" xorm:"text notnull 'title'"`
	Content    string     `json:"content" xorm:"longtext notnull 'content'"`
	Views      uint64     `json:"article_views" xorm:"bigint(20) notnull default(0) 'article_views'"`
	Status     int        `json:"article_status,omitempty" xorm:"tinyint(4) notnull default(0) 'article_status'"`
	CreateTime *time.Time `json:"create_time,omitempty" xorm:"datetime created notnull 'create_time'"`
	UpdateTime *time.Time `json:"update_time,omitempty" xorm:"datetime updated notnull 'update_time'"`
	DeleteTime *time.Time `json:"delete_time,omitempty" xorm:"datetime 'delete_time'"`
}

// ArticlesByPage handles articles by page number.
func ArticlesByPage(limit int, page int) []Article {
	db := orm.NewSession()
	defer db.Close()

	var articles []Article
	err := db.Table("article").
		Cols("ID", "slug", "title", "content", "article_views", "create_time").
		Where("article_status = 1").
		Limit(limit, limit*(page-1)).
		Desc("create_time").
		Find(&articles)
	if err != nil {
		log.Errorf("ArticlesByPage throw error: %s", err)
	}

	return articles
}

// AddArticle handles and add article.
func AddArticle(article *Article) bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.InsertOne(article)
	if err != nil {
		log.Errorf("AddArticle throw error: %s", err)
	}

	return affected > 0
}

// ArticleExist if article exist
func ArticleExist(id uint64) bool {
	db := orm.NewSession()
	defer db.Close()

	has, err := db.Exist(&Article{
		ID: id,
	})
	if err != nil {
		log.Errorf("ArticleExist throw error: %s", err)
	}

	return has
}

// ArticlesCount returns count number.
func ArticlesCount() int {
	db := orm.NewSession()
	defer db.Close()

	article := &Article{}

	count, err := db.Count(article)
	if err != nil {
		log.Errorf("ArticlesCount throw error: %s", err)
	}

	return int(count)
}

// ArticleByID returns article by ID.
func ArticleByID(id uint64) (*Article, bool) {
	db := orm.NewSession()
	defer db.Close()

	article := &Article{
		ID: id,
	}

	has, err := db.Get(article)
	if err != nil {
		log.Errorf("ArticleByID throw error: %s", err)
	}

	return article, has
}

// ArticleBySlug returns article by slug.
func ArticleBySlug(slug string) (*Article, bool) {
	db := orm.NewSession()
	defer db.Close()

	article := &Article{
		Slug: slug,
	}

	has, err := db.Get(article)
	if err != nil {
		log.Errorf("ArticleBySlug throw error: %s", err)
	}

	return article, has
}

// UpdateByID updates article data by ID.
func (a *Article) UpdateByID() bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.ID(a.ID).Update(a)
	if err != nil {
		log.Errorf("Article update throw error: %s", err)
	}

	return affected > 0
}
