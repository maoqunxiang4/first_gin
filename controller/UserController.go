package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goPro/common"
	"goPro/model"
	"goPro/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB = common.GetDB("t_user")

func Register() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")
		password := ctx.Query("password")
		phone := ctx.Query("phone")

		if len(phone) != 11 {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
			return
		}

		if len(password) <= 6 {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
			return
		}

		var user model.User
		db.Table("t_user").Where(" phone = ? ", phone).First(&user)
		if user.ID != 0 {
			response.Response(ctx, http.StatusInternalServerError, 500, nil, "There is a same user now")
			return
		}

		//对密码进行加密
		headPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
			return
		}
		newUser := model.User{
			UserName: username,
			Password: string(headPassword),
			Phone:    phone,
		}
		db.Table("t_user").Create(&newUser)

		response.Success(ctx, gin.H{}, "")
	}
}

func Login() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//在进行数据库的操作事前，一定要使用db.Table("t_user")来指定查找的表对象
		db = db.Table("t_user")
		phone := ctx.Query("phone")
		password := ctx.Query("password")

		var user model.User
		db.Where("phone = ?", phone).First(&user)
		if user.ID == 0 {
			response.Response(ctx, http.StatusBadRequest, 400, gin.H{}, "Your phone is not registe ")
			return
		}

		fmt.Println(user)

		//对比两次密码是否正确,直接通过输入的password与数据库中的Password进行比较
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			fmt.Println(err)
			response.Response(ctx, http.StatusBadRequest, 400, gin.H{}, "the password not right ")
			return
		}

		token, err := common.ReleaseToken(user)
		if err != nil {
			response.Response(ctx, http.StatusInternalServerError, 500, gin.H{}, "token is fail")
			return
		}

		response.Success(ctx, gin.H{"token": token}, "")
	}
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": user}, "")
}
