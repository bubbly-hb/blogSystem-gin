# select 无效(关键字加反引号)
```go
err = db.Preload("Category").Select("article.id, title, desc, category.name, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
```
改为：
```go
err = db.Select("article.id, title, `desc`, category.name, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("Category").Find(&articles).Error
```
