/**
 * @Author: Robby
 * @File name: sso.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package controllers

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
	"sso/model"
	"sso/oauth2"
	"sso/service"
	"time"

	"sso/session"
	"sso/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/*
认证的调用顺序：
AuthorizeHandler.....................
LoginHandler GET.....................
LoginHandler POST.....................
ReAuthorizeHandler.....................
*/

// AuthorizeHandler Get接口，第一次调用
func AuthorizeHandler(c *gin.Context) {

	if err := session.Delete(c.Writer, c.Request, "RequestForm"); err != nil {
		zap.L().Error("[AuthorizeHandler]：session.Delete", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	if err := oauth2.Srv.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
		zap.L().Error("[AuthorizeHandler]：Srv.HandleAuthorizeRequest", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

}

// ReAuthorizeHandler Get接口，第二次调用，数据从session中取出来
func ReAuthorizeHandler(c *gin.Context) {
	var err error
	var requestFormString string
	var requestForm url.Values

	if requestFormString, err = session.Get(c.Request, "RequestForm"); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：session.Get", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	if requestForm, err = url.ParseQuery(requestFormString); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：url.ParseQuery", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	// 给请求的form赋值
	c.Request.Form = requestForm

	if err = session.Delete(c.Writer, c.Request, "RequestForm"); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：session.Delete", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	// 这里由于没有在浏览器写入session_id，导致无法获取到RequestForm对应的数据
	if err = oauth2.Srv.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：oauth2.Srv.HandleAuthorizeRequest", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

}

// 获取requestForm的数据, 为LoginHandler服务
func getRequestForm(c *gin.Context) (data *model.ClientScope, err error) {

	requestForm, err := session.Get(c.Request, "RequestForm")
	if err != nil {
		zap.L().Error("[LoginHandler]：session.Get", zap.Error(err))
		return nil, errors.New(CodeBadRequest.ToString())
	}

	if requestForm == "" {
		zap.L().Info("[LoginHandler]：requestForm == '' ")
		return nil, errors.New(CodeBadRequest.ToString())
	}

	decodeForm, err := url.ParseQuery(requestForm)
	if err != nil {
		return nil, errors.New(CodeServerInternalError.ToString())
	}

	// Get client_id and scope_name from user
	clientID := decodeForm.Get("client_id")
	scope := decodeForm.Get("scope")
	clientObj := utils.GetClientObj(clientID)
	scopeObj := utils.GetClientScope(clientID, scope)
	if scopeObj == nil {
		zap.L().Error("[LoginHandler]：bad scope")
		return nil, errors.New(CodeInvalidParam.ToString())
	}

	return &model.ClientScope{
		Client: clientObj,
		Scope:  scopeObj,
	}, nil

}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {

	switch c.Request.Method {

	case http.MethodGet:
		// 获取scope，没有意义
		//data, err := getRequestForm(c)

		//if err != nil {
		//	code, err := strconv.Atoi(err.Error())
		//	if err != nil {
		//		zap.L().Error("[LoginHandler]：strconv.Atoi", zap.Error(err))
		//		ResponseError(c, CodeServerInternalError)
		//	}
		//	ResponseError(c, ResCode(code))
		//	return
		//}

		//tmpl, err := template.ParseFiles("tpl/login.html")
		//if err != nil {
		//	zap.L().Error("[LoginHandler]：template parse error")
		//	ResponseError(c, CodeServerInternalError)
		//	return
		//}
		//
		////if err = tmpl.Execute(c.Writer, data); err != nil {
		//if err = tmpl.Execute(c.Writer, nil); err != nil {
		//	zap.L().Error("[LoginHandler]：template execute error")
		//}

		// 让vue来显示登录页面

		ResponseToLogin(c, "")


	case http.MethodPost:

		// 获取用户提交的用户名和密码
		params := new(model.UserLoginParam)
		if err := c.ShouldBindJSON(params); err != nil {
			zap.L().Error("登录参数校验错误", zap.Error(err))
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				ResponseError(c, CodeInvalidParam)
				return
			}

			ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans)))
			return
		}

		// 基于用户名密码获取userid
		userID, err := service.GetUserIdByNamePwd(params)

		if err != nil || userID == "" {
			zap.L().Error("[LoginHandler]：service.GetUserIdByNamePwd", zap.Error(err))
			ResponseErrorWithMsg(c, CodeInvalidPassword, err.Error())
		}

		// 将userid写入到session中
		if err = session.Set(c.Writer, c.Request, "LoggedInUserID", userID); err != nil {
			zap.L().Error("[LoginHandler]：session.Set", zap.Error(err))
			ResponseError(c, CodeServerInternalError)
			return
		}

		c.Redirect(http.StatusFound, "/api/v1/reauthorize")
		return

		//// csrf token verify
		//if c.PostForm("type") == "password" {
		//
		//	// 如果传递的是空密码那么binding就会校验，依托的是github.com/gin-gonic/gin/binding库，自动校验参数返回
		//	userLoginParam := &model.UserLoginParam{
		//		Username: c.PostForm("username"),
		//		Password: c.PostForm("password"),
		//	}
		//
		//	userID, err := service.GetUserIdByNamePwd(userLoginParam)
		//
		//	if err != nil {
		//		zap.L().Error("[LoginHandler]：service.GetUserIdByNamePwd", zap.Error(err))
		//		ResponseErrorWithMsg(c, CodeInvalidPassword, err.Error())
		//	}
		//
		//	if userID == "" {
		//
		//		tmpl, err := template.ParseFiles("sso/tpl/login.html")
		//		if err != nil {
		//			zap.L().Error("[LoginHandler]：template.ParseFiles", zap.Error(err))
		//			ResponseError(c, CodeServerInternalError)
		//			return
		//		}
		//
		//		data, err := getRequestForm(c)
		//
		//		if err != nil {
		//			code, err := strconv.Atoi(err.Error())
		//			if err != nil {
		//				zap.L().Error("[LoginHandler]：strconv.Atoi", zap.Error(err))
		//				ResponseError(c, CodeServerInternalError)
		//			}
		//			ResponseError(c, ResCode(code))
		//		}
		//
		//		if data != nil {
		//			data.Error = "用户名或密码错误"
		//		}
		//
		//		if err = tmpl.Execute(c.Writer, data); err != nil {
		//			zap.L().Error("[LoginHandler]：tmpl.Execute")
		//		}
		//
		//	}
		//
		//	// 将userid写入到session中
		//	if err = session.Set(c.Writer, c.Request, "LoggedInUserID", userID); err != nil {
		//		zap.L().Error("[LoginHandler]：session.Set", zap.Error(err))
		//		ResponseError(c, CodeServerInternalError)
		//		return
		//	}
		//
		//	c.Redirect(http.StatusFound, "/api/v1/reauthorize")
		//	return
		//
		//}
		//
		//// csrf token error
		//ResponseError(c, CodeBadRequest)

	}
}

// LogoutHandler 登出
func LogoutHandler(c *gin.Context) {

	var redirectUri string

	if redirectUri = c.Query("redirect_uri"); redirectUri == "" {
		zap.L().Error("[LogoutHandler]：c.Query", zap.Error(errors.New("No RedirectUri")))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := session.Delete(c.Writer, c.Request, "LoggedInUserID"); err != nil {
		zap.L().Error("[LogoutHandler]：session.Delete", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	c.Redirect(http.StatusFound, redirectUri)

}

// TokenHandler 获取token或刷新token
func TokenHandler(c *gin.Context) {

	if err := oauth2.Srv.HandleTokenRequest(c.Writer, c.Request); err != nil {
		zap.L().Error("[TokenHandler]：oauth2.Srv.HandleTokenRequest", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

}

// VerifyHandler 验证token
func VerifyHandler(c *gin.Context) {

	token, err := oauth2.Srv.ValidationBearerToken(c.Request)
	if err != nil {
		zap.L().Error("[VerifyHandler]：oauth2.Srv.ValidationBearerToken", zap.Error(err))
		ResponseError(c, CodeInvalidToken)
		return
	}

	clientInfo, err := oauth2.Manager.GetClient(context.Background(), token.GetClientID())
	if err != nil {
		zap.L().Error("[VerifyHandler]：oauth2.Manager.GetClient", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"user_id":    token.GetUserID(),
		"client_id":  token.GetClientID(),
		"scope":      token.GetScope(),
		"domain":     clientInfo.GetDomain(),
	}

	ResponseSuccess(c, data)

}
