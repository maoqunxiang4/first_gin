package middleware

import (
	"github.com/gin-gonic/gin"
	"goPro/common"
	"goPro/model"
	"net/http"
	"strings"
)

// 如果token存服务端的话可以不用查库，但是如果服务端不存token，而是前端存token，那么
// 控制token有效性的方式就只能通过之前生成时设置的时间了，这样在这期间如果用户被删除的话
// 这个token还是有效的话就是Bug了，所以如果是存在前端的话就要再查一次库
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		//在postman中，我们通过authorization来向接口传值
		tokenString := ctx.GetHeader("authorization")

		//validate token formate
		//这里对我们获取的token由于是通过authorization进行传值的，所以他的前缀一定是“Bearer ”
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后获取claimz中的userId
		userId := claims.UserId
		db := common.GetDB("t_user")
		var user model.User
		db.First(&user, userId)

		//用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401, "msg": "权限不足",
			})
			ctx.Abort()
			return
		}

		//用户存在，将user的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
