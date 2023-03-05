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

// 查询分类下的所有文章
func GetArticlesByCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid, _ := strconv.Atoi(c.Query("cid"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	articles, total, errCode := model.GetArticlesByCategory(cid, pageSize, pageNum)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"data":    articles,
			"total":   total,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}

// 查询单个文章
func GetArticle(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("id"))
	article, errorCode := model.GetArticle(articleId)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errorCode,
			"data":    article,
			"message": errmsg.GetErrMsg(errorCode),
		},
	)
}

// 查询文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	articles, total, errCode := model.GetArticles(title, pageSize, pageNum)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"data":    articles,
			"total":   total,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}

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
