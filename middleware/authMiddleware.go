package middleware

import (
	"net/http"
	"strings"

	"github.com/bubbly-hb/blogSystem-gin-vue/common"
	"github.com/bubbly-hb/blogSystem-gin-vue/dao"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "auth check failed"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "auth check failed"})
			ctx.Abort()
			return
		}
		userID := claims.UserID // 验证通过后获取claims中的userID
		user, err := dao.GetUserByID(userID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "auth check failed"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user) // 用户存在，将用户信息写入context
		ctx.Next()
	}
}
