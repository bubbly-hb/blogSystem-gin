package controller

import (
	"log"
	"net/http"

	"github.com/bubbly-hb/blogSystem-gin-vue/common"
	"github.com/bubbly-hb/blogSystem-gin-vue/dao"
	"github.com/bubbly-hb/blogSystem-gin-vue/model"
	"github.com/bubbly-hb/blogSystem-gin-vue/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "bcrypt failed",
		})
		return
	}
	user := &model.User{
		Name:     name,
		Password: string(hashedPassword),
		Email:    email,
	}
	err = dao.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "create user failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Register Success",
	})
}

func Login(ctx *gin.Context) {
	email := ctx.PostForm("email") // 获取参数
	password := ctx.PostForm("password")
	if len(email) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "email can not be nil",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "password can not be less than 6 digits",
		})
		return
	}
	user, err := dao.GetUserByEmail(email) // 判断用户是否存在
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "user not exist",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) // 校验密码是否正确
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Incorrect email or password",
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "release token failed",
		})
		log.Println("token generate error: ", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{ // 返回结果
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "Login Success",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{ // 返回结果
		"code": 200,
		"data": gin.H{"user": user},
	})
}
