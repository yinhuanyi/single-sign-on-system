package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"user-server/controllers/response"
	"user-server/model"
	"user-server/service"
)

// SSOLoginHandler 前端登录请求
func SSOLoginHandler(c *gin.Context)  {
	fmt.Println("c..................")
	response.ResponseSuccess(c, nil)
}

// UserProfileHandler 设置
func UserProfileHandler(c *gin.Context)  {

	// 获取userId
	userIdValue, isExists := c.Get("userId")
	if  !isExists {
		zap.L().Error("Can Not Get UserId")
		response.ResponseErrorWithMsg(c, response.CodeUserNotExist, "Can Not Get UserId")
	}
	userIdString := userIdValue.(string)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		zap.L().Error("userId assert error")
		response.ResponseError(c, response.CodeServerInternalError)
	}

	// 打印userId日志
	zap.L().Info("Get userId Success! ", zap.Int("userId", userId))

	// 基于userId获取到用户的info信息
	user := &model.User{UserId: userId}
	userInfo, err := service.GetUser(user)
	if err != nil {
		zap.L().Error("Get UserInfo Error", zap.Error(err))
		response.ResponseError(c, response.CodeServerBusy)
	}

	response.ResponseSuccess(c, userInfo)
}


