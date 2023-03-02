package model

import "github.com/bubbly-hb/blogSystem-gin/utils/errmsg"

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategory(name string) (cid uint, code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return category.ID, errmsg.ERROR_CATEGORYNAME_USED
	}
	return 0, errmsg.SUCCESS
}

func CreateCategory(category *Category) int {
	err := db.Create(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCategorys(categoryName string, pageSize, pageNum int) ([]Category, int) {
	var categorys []Category
	if categoryName != "" {
		db.Select("id, name").Where(
			"name LIKE ?", categoryName+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categorys)
	} else {
		db.Select("id, name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categorys)
	}
	return categorys, len(categorys)
}

func UpdateCategory(id uint, category Category) int {
	err := db.Model(&category).Where("id = ?", id).Updates(map[string]interface{}{"name": category.Name}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCategory(id uint) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
