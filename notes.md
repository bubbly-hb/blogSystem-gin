# select 无效(关键字加反引号)
```go
err = db.Preload("Category").Select("article.id, title, desc, category.name, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
```
改为：
```go
err = db.Select("article.id, title, `desc`, category.name, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("Category").Find(&articles).Error
```

# (SCC-ST1012) Poorly chosen name for error variable
It is recommended that the error variables that are part of an API should be named as errFoo, ErrSomethingBad, ErrKindFoo or BazError.
```go
var (
	ErrTokenExpired     = errors.New("token已过期,请重新登录")
	ErrTokenNotValidYet = errors.New("token无效,请重新登录")
	ErrTokenMalformed   = errors.New("token不正确,请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个token,请重新登录")
)
```
而不是
```go
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)
```