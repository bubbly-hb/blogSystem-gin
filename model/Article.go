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

// func GetUsers(userName string, pageSize, pageNum int) ([]User, int) {
// 	var users []User
// 	if userName != "" {
// 		db.Select("id, username, role, created_at").Where(
// 			"username LIKE ?", userName+"%",
// 		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
// 	} else {
// 		db.Select("id, username, role, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
// 	}
// 	return users, len(users)
// }

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
