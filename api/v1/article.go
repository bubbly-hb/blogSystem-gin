package v1

import (
	"net/http"
	"strconv"

	"github.com/bubbly-hb/blogSystem-gin/model"
	"github.com/bubbly-hb/blogSystem-gin/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)

	errorCode := model.CreateArticle(&article)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errorCode,
			"data":    article,
			"message": errmsg.GetErrMsg(errorCode),
		},
	)
}

// // 查询用户列表
// func GetUsers(c *gin.Context) {
// 	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
// 	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
// 	username := c.Query("username")

// 	switch {
// 	case pageSize >= 100:
// 		pageSize = 100
// 	case pageSize <= 0:
// 		pageSize = 10
// 	}

// 	if pageNum == 0 {
// 		pageNum = 1
// 	}

// 	users, total := model.GetUsers(username, pageSize, pageNum)

// 	code := errmsg.SUCCESS
// 	c.JSON(
// 		http.StatusOK, gin.H{
// 			"status":  code,
// 			"data":    users,
// 			"total":   total,
// 			"message": errmsg.GetErrMsg(code),
// 		},
// 	)
// }

// 编辑文章
func UpdateArticle(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("id"))
	var article model.Article
	c.ShouldBindJSON(&article) // 这里注意加&

	errCode := model.UpdateArticle(uint(aid), article)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	aid, _ := strconv.Atoi(c.Param("id"))
	errCode := model.DeleteArticle(uint(aid))
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}
