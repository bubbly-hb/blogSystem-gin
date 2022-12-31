package controller

import (
	"log"
	"math/rand"
	"net/http"
	"time"

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
		name = RandomString(10)
	}
	log.Println(name, password, email)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Register Success",
	})
}

func RandomString(n int) string {
	data := []byte{}
	for i := 0; i < 26; i++ {
		data = append(data, byte('a'+i))
		data = append(data, byte('A'+i))
	}
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		res[i] = data[rand.Intn(len(data))]
	}
	return string(res)
}
