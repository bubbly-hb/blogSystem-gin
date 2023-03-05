package model

import (
	"github.com/bubbly-hb/blogSystem-gin/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

func CreateArticle(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取分类下的所有文章
func GetArticlesByCategory(cid, pageSize, pageNum int) ([]Article, int, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&articles).Error
	if err != nil {
		return articles, len(articles), errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	return articles, len(articles), errmsg.SUCCESS
}

// 获取单篇文章
func GetArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where(
		"id = ?", id).First(&article).Error
	db.Model(&article).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// 获取文章列表
func GetArticles(title string, pageSize, pageNum int) ([]Article, int, int) {
	var articles []Article
	var err error
	if title != "" {
		err = db.Select("article.id, title, `desc`, category.name, created_at").Limit(pageSize).Offset((pageNum-1)*pageSize).Joins("Category").Where("title like ?", title+"%").Find(&articles).Error
	} else {
		err = db.Select("article.id, title, `desc`, category.name, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("Category").Find(&articles).Error
	}
	if err != nil {
		return articles, len(articles), errmsg.ERROR
	}
	return articles, len(articles), errmsg.SUCCESS
}

func UpdateArticle(aid uint, article Article) int {
	dic := map[string]interface{}{
		"title":   article.Title,
		"cid":     article.Cid,
		"desc":    article.Desc,
		"content": article.Content,
		"img":     article.Img,
	}
	err := db.Model(&article).Where("id = ?", aid).Updates(dic).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticle(aid uint) int {
	err := db.Where("id = ?", aid).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
