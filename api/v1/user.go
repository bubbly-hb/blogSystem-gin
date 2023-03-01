package v1

import (
	"net/http"
	"strconv"

	"github.com/bubbly-hb/blogSystem-gin/model"
	"github.com/bubbly-hb/blogSystem-gin/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	// validator TODO
	// ...
	// 密码加密存储 scrypt TODO
	// ...
	errorCode := model.CheckUser(user.Username)
	if errorCode == errmsg.SUCCESS {
		model.CreateUser(&user)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errorCode,
			"message": errmsg.GetErrMsg(errorCode),
		},
	)
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	users, total := model.GetUsers(username, pageSize, pageNum)

	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    users,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DeleteUser(c *gin.Context) {

}
