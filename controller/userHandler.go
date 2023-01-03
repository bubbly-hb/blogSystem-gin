package controller

import (
	"net/http"

	"github.com/bubbly-hb/blogSystem-gin-vue/dao"
	"github.com/bubbly-hb/blogSystem-gin-vue/model"
	"github.com/bubbly-hb/blogSystem-gin-vue/util"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	if len(email) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "email nil",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "password cannot be less than 6 numbers",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	isEmailExist := dao.IsEmailExist(email)
	if isEmailExist {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "email already exist",
		})
		return
	}

	user := &model.User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	err := dao.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "create user failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Register Success",
	})
}
