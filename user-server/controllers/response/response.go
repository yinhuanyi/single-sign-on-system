/**
 * @Author: Robby
 * @File name: response.go
 * @Create date: 2021-11-04
 * @Function: 封装业务状态码
 **/

package response

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResCode int64


const (
	CodeSuccess ResCode = 20000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeBadRequest
	CodeNeedLogin
	CodeInvalidToken
	CodeServerBusy
	CodeServerInternalError

	//CodeBookExist
	//CodeBookNotExist

	//CodeTokenNeeded ResCode = 400
	//CodeTokenExpired ResCode = 400
	//CodeTokenInvalid ResCode = 400
	//CodeTokenNotActived ResCode = 400
	//CodeTokenUnrecognized ResCode = 400
	//CodeRequestBusy ResCode = 400

	CodeRedirect ResCode = 20302
	CodeFlushToken ResCode = 20303

)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:             "请求成功",
	CodeInvalidParam:        "请求参数错误",
	CodeUserExist:           "用户已经存在",
	CodeUserNotExist:        "用户不存在",
	CodeInvalidPassword:     "密码错误",
	CodeBadRequest:          "请求错误",
	CodeNeedLogin:           "需要登录",
	CodeInvalidToken:        "无效的token",
	CodeServerBusy:          "服务器繁忙",
	CodeServerInternalError: "服务内部错误",
	//CodeBookExist:			 "电子书已经存在",
	//CodeBookNotExist:		 "电子书不存在",

	//CodeTokenNeeded:       "用户未登录",
	//CodeTokenInvalid:      "Token不合法",
	//CodeTokenExpired:      "Token过期",
	//CodeTokenNotActived:   "Token未激活",
	//CodeTokenUnrecognized: "Token无法识别",
	//
	//CodeRequestBusy: "请求过于频繁, 请稍后重试",
	CodeRedirect:           "302重定向请求",
	CodeFlushToken: 		"删除浏览器token",

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

// ResponseHttpError 这是自定义的HTTP Error
func ResponseHttpError(c *gin.Context, httpCode ResCode)  {
	rd := &ResponseData{
		Msg:  httpCode.GetMsg(),
	}

	c.JSON(int(httpCode), rd)
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

// Response302 让前端更新access_token和refresh_token
func Response302(c *gin.Context, data interface{})  {
	rd := &ResponseData{
		Success: true,
		Code: CodeRedirect,
		Msg:  CodeRedirect.GetMsg(),
		Data: data,
	}

	c.JSON(http.StatusOK, rd)
}

// Response303 让前端删除access_token和refresh_token
func Response303(c *gin.Context,)  {
	rd := &ResponseData{
		Success: true,
		Code: CodeFlushToken,
		Msg:  CodeFlushToken.GetMsg(),
	}

	c.JSON(http.StatusOK, rd)
}