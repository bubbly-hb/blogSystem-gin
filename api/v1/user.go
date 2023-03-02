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
	_, errorCode := model.CheckUser(user.Username)
	if errorCode == errmsg.SUCCESS {
		model.CreateUser(&user)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errorCode,
			"data":    user,
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
func UpdateUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	c.ShouldBindJSON(&user) // 这里注意加&
	uidByName, errCode := model.CheckUser(user.Username)
	// fmt.Println(uidByName, errCode, uid, user.Username, user.Role)
	if errCode == errmsg.SUCCESS || uidByName == uint(uid) {
		errCode = errmsg.SUCCESS
		model.UpdateUser(uint(uid), user)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))
	errCode := model.DeleteUser(uint(uid))
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		},
	)
}
