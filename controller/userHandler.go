package controller

import (
	"log"
	"net/http"

	"github.com/bubbly-hb/blogSystem-gin/common"
	"github.com/bubbly-hb/blogSystem-gin/dao"
	"github.com/bubbly-hb/blogSystem-gin/dto"
	"github.com/bubbly-hb/blogSystem-gin/model"
	"github.com/bubbly-hb/blogSystem-gin/response"
	"github.com/bubbly-hb/blogSystem-gin/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")

	if len(email) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "email can not be nil")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "password can not be less than 6 numbers")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	isEmailExist := dao.IsEmailExist(email)
	if isEmailExist {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "email already exist")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "bcrypt failed")
		return
	}
	user := &model.User{
		Name:     name,
		Password: string(hashedPassword),
		Email:    email,
	}
	err = dao.CreateUser(user)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "create user failed")
		return
	}

	response.Succces(ctx, nil, "Register Success")
}

func Login(ctx *gin.Context) {
	email := ctx.PostForm("email") // 获取参数
	password := ctx.PostForm("password")
	if len(email) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "email can not be nil")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "password can not be less than 6 numbers")
		return
	}
	user, err := dao.GetUserByEmail(email) // 判断用户是否存在
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "user not exist")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) // 校验密码是否正确
	if err != nil {
		response.Fail(ctx, nil, "Incorrect email or password")
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "release token failed")
		log.Println("token generate error: ", err)
		return
	}
	response.Succces(ctx, gin.H{"token": token}, "Login Success")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Succces(ctx, gin.H{"user": dto.ToUserDto(user.(*model.User))}, "get info success")
}
