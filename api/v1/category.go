package v1

import (
	"net/http"
	"strconv"

	"github.com/bubbly-hb/blogSystem-gin/model"
	"github.com/bubbly-hb/blogSystem-gin/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// 添加文章类别
func AddCategory(c *gin.Context) {
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	// validator TODO
	// ...
	_, errorCode := model.CheckCategory(category.Name)
	if errorCode == errmsg.SUCCESS {
		model.CreateCategory(&category)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errorCode,
			"message": errmsg.GetErrMsg(errorCode),
		},
	)
}

// 查询类别列表
func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	name := c.Query("name")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	categorys, total := model.GetCategorys(name, pageSize, pageNum)

	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    categorys,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 编辑类别
func UpdateCategory(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	var category model.Category
	c.ShouldBindJSON(&category) // 这里注意加&
	cidByName, errCode := model.CheckCategory(category.Name)
	if errCode == errmsg.SUCCESS || cidByName == uint(cid) {
		errCode = errmsg.SUCCESS
		model.UpdateCategory(uint(cid), category)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}

// 删除类别
func DeleteCategory(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("id"))
	errCode := model.DeleteCategory(uint(cid))
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}
