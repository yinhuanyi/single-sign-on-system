package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user-server/controllers/response"
)

// SSOLoginHandler 前端登录请求
func SSOLoginHandler(c *gin.Context)  {
	fmt.Println("c..................")
	response.ResponseSuccess(c, nil)
}

// UserProfileHandler 设置
func UserProfileHandler(c *gin.Context)  {

	userId, isExists := c.Get("userId")
	if isExists {
		fmt.Printf("userId=%s\n", userId)
	}

	res := map[string]interface{}{
		"uid": "1020020020",
		"username": "robby",
		"age": 30,
	}
	response.ResponseSuccess(c, res)
}


