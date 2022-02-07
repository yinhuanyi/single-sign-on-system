/**
 * @Author: Robby
 * @File name: response.go
 * @Create date: 2021-11-04
 * @Function: 封装业务状态码
 **/

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeBadRequest
	CodeNeedLogin
	CodeInvalidToken
	CodeServerBusy
	CodeServerInternalError

	CodeFrontLogin
	CodeUsernameIsBlocked
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:             "Success",
	CodeInvalidParam:        "请求参数错误",
	CodeUserExist:           "用户已经存在",
	CodeUserNotExist:        "用户不存在",
	CodeInvalidPassword:     "密码错误",
	CodeBadRequest:          "请求错误",
	CodeNeedLogin:           "需要登录",
	CodeInvalidToken:        "无效的token",
	CodeServerBusy:          "服务器繁忙",
	CodeServerInternalError: "服务内部错误",

	CodeFrontLogin:          "Login",
	CodeUsernameIsBlocked:	 "当前用户被锁定，请5分钟后重试",
}

type ResponseData struct {
	Success bool `json:"success"`
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (c ResCode) GetMsg() string {
	msg, ok := CodeMsgMap[c]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}

func (c ResCode) ToString() string {
	return strconv.Itoa(int(c))
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Success: false,
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	}

	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Success: false,
		Code: code,
		Msg:  fmt.Sprintf("%s: %v", code.GetMsg(), msg),
		Data: nil,
	}

	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Success: true,
		Code: CodeSuccess,
		Msg:  CodeSuccess.GetMsg(),
		Data: data,
	}

	c.JSON(http.StatusOK, rd)

}


func ResponseToLogin(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Success: true,
		Code: CodeFrontLogin,
		Msg:  CodeFrontLogin.GetMsg(),
		Data: data,
	}

	c.JSON(http.StatusOK, rd)

}
